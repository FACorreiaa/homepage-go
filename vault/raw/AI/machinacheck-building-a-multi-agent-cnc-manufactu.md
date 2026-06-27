---
source: https://huggingface.co/blog/machinacheck-building-a-multi-agent-cnc-manufactu
ingested_at: 2026-05-11T09:24:55.948043Z
type: web
status: uncompiled
topic: AI
---
Hugging Face
Buckets
NEW
MachinaCheck: Building a Multi-Agent CNC Manufacturability System on AMD MI300X
Team
Article
Published May 10, 2026
Syed Muhammad Sarmad
sarmaddev
Follow
Built at the AMD Developer Hackathon on lablab.ai — May 2026
The Problem We Solved
Walk into any small CNC machine shop and ask the manager how they decide whether to accept a customer job.
The answer is almost always the same: they print the drawing, read every dimension by hand, walk around the shop checking which tools are available, estimate whether their machines can hold the required tolerances, and write notes on a clipboard. The whole process takes 30 to 60 minutes per drawing. For a busy shop receiving 10 to 20 RFQs per week, that is 5 to 20 hours of skilled manager time spent on feasibility analysis alone.
Sometimes they get it wrong. They accept a job, start production, and discover halfway through that they don't have the right tap or that their mill cannot hold the tolerance on a critical feature. The part gets scrapped. The customer is unhappy. The machine time is lost.
We built MachinaCheck to eliminate this problem entirely.
What MachinaCheck Does
MachinaCheck is a multi-agent AI system. You upload a STEP file — the standard CAD format that customers send to machine shops — along with three simple inputs: material type, required tolerance, and any thread specifications. Thirty seconds later you have a complete manufacturability report telling you exactly whether you can make the part, what tools you need, what is missing, and what actions to take before starting production.
No manual drawing reading. No walking around the shop. No guesswork.
Why We Built It on AMD MI300X
Before explaining the architecture, this point deserves its own section because it is not just a technical choice — it is a business requirement.
Manufacturing customers sign NDAs. Their STEP files contain proprietary geometry representing years of engineering work and millions of dollars in R&D. The hole pattern on a medical device or the pocket geometry on an aerospace component is confidential intellectual property.
Sending that data to OpenAI, Anthropic, or any commercial API endpoint is a confidentiality violation. Full stop.
The AMD Instinct MI300X changes this equation completely. With 192GB of HBM3 VRAM and 5.3 TB/s of memory bandwidth, we run Qwen 2.5 7B Instruct entirely on-premise. No data leaves the shop's infrastructure. No STEP geometry is transmitted to a third-party server. The customer's IP stays where it belongs.
This is what \
The Agent Architecture
MachinaCheck uses a five-component pipeline built with LangChain and orchestrated via FastAPI.
Component 1 — STEP File Parser (Pure Python, No LLM)
We use cadquery, a Python library built on OpenCASCADE, to parse STEP files directly. This gives us mathematically exact feature extraction:
All cylindrical holes with diameter and depth
Flat surfaces and their areas
Chamfers and fillets
Bounding box dimensions
Total volume and surface area
This extraction is 100% accurate because it reads the mathematical geometry directly — no vision model, no OCR, no approximation. A Ø6.0mm hole is exactly Ø6.0mm in the output.
def extract_features(step_file_path: str) -> dict:\n    model = cq.importers.importStep(step_file_path)\n    shape = model.val()\n    bb = shape.BoundingBox()\n    \n    holes = {}\n    for face in model.faces().vals():\n        adaptor = BRepAdaptor_Surface(face.wrapped)\n        if adaptor.GetType() == GeomAbs_Cylinder:\n            radius = adaptor.Cylinder().Radius()\n            diameter = round(radius * 2, 3)\n            holes[diameter] = holes.get(diameter, 0) + 1\n    \n    return {\n        \
Agent 1 — Operations Classifier (Qwen 2.5 7B)
The extracted geometry plus user inputs — material, tolerance, threads — are passed to Qwen 2.5 7B running on AMD MI300X via vLLM.