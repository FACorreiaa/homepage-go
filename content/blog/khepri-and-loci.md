---
title: "Khepri and Loci are live"
summary: "Two products I have been building in private are now on the public internet: an AI coach that remembers, and a travel planner that keeps the trip."
category: Building in Public
date: "2026-09-07"
---

Two products I have been building in private are now on the public internet.

[Khepri](https://kheprios.com) is an AI coach with one memory. [Loci](https://lociai.fyi) is a travel planner that produces a trip you can keep. Both have case studies on this site: [/projects/khepri](/projects/khepri) and [/projects/loci](/projects/loci).

## Khepri: the assistant that does not reset

Most AI tools are a blank slate every time you open them. That is fine for a one-off question. It is useless if you are training for a hike, recovering from a bad back, or trying to stay consistent for more than a week.

Khepri is the bet the other way: one coach, one memory, still useful in month eight. It holds goals, training, meals, sleep, documents, and check-ins, then answers from that record instead of the last ten messages. The same thread is on the web, in Telegram, and inside any client that speaks MCP.

The stack is the one I already operate: Go, templ, HTMX, Alpine, PostgreSQL. You can bring your own OpenAI, Anthropic, Gemini, or xAI key. If you do not, there is a working floor so the coach is not a blank error page.

It is free to try. No card. The product name is Khepri; the repo is still called north-web-app. That is a rename I have not done, not a second product.

## Loci: a plan that survives the chat

The other failure mode I kept hitting is travel. You ask a chatbot for a weekend, you get a wall of text, and by the time you are in the city the plan is buried in Notes.

Loci takes a city and a mood and streams back an itinerary of real places: mapped, ordered, editable. You save it, export it, and open it again when you are standing there. The two-city compare is for the specific Iberia-weekend decision (Évora or Beja from Porto this Saturday?). The MCP server is so Claude, or whatever you already run, can build and save trips against the same product.

The stack is Go and Connect RPC on the API, PostgreSQL with PostGIS, and a SolidStart client. Live weather, holidays, and trip-kit export are in. Offline trips and a 24/7 personal agent are not, and I am not going to write as if they are.

The public URL is [lociai.fyi](https://lociai.fyi). loci.app is still a parked name. I am listing the site people can actually open.

## What this is not

No user counts. No "quietly in use at" line. Khepri and Loci sit next to Norviq on [/projects](/projects) because they are hosted products I shipped and still stand behind, not because a metric made them look finished.

If you want a weekend plotted or a coach that remembers March, the doors are open.
