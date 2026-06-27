---
source: Stack Overflow RSS Feed
ingested_at: 2026-05-11T08:06:16.066960
type: web
status: uncompiled
---

# Unknown Error code 1032 from Apple Intelligence

**Stack Overflow Question**

**Link:** https://stackoverflow.com/questions/79939126/unknown-error-code-1032-from-apple-intelligence

**Published:** 2026-05-11T03:01:51Z

**Updated:** 2026-05-11T03:01:51Z

**Author:** Jon Vogel

**Tags:** swift, xcode, apple-intelligence, foundation-models

**Summary:**

            <p>I'm using Apple intelligence to power suggestions to user of my app. It's a great feature but I'm baffled about this error.</p>
<pre class="lang-bash prettyprint-override"><code>Domain=FoundationModels.LanguageModelSession.GenerationError Code=-1 &quot;(null)&quot; UserInfo={NSMultipleUnderlyingErrorsKey=(
    &quot;Error Domain=ModelManagerServices.ModelManagerError Code=1032 \&quot;(null)\&quot; UserInfo={NSMultipleUnderlyingErrorsKey=(\n)}&quot;
)} 
</code></pre>
<p>My code uses the &quot;stream&quot; feature and <code>@Generable</code> structures to give the user a great experience.</p>
<pre class="lang-swift prettyprint-override"><code>    func generateIdeas(name: String, hero: String, creativity: Creativity, numberOfIdeas: Int) throws -&gt; LanguageModelSession.ResponseStream&lt;[ColorChip]&gt; {

        try self.checkAvailability()
        
        let prompt = &quot;&quot;&quot;
        Please generate at least \(numberOfIdeas) different and unique primary colors for the users idea which they are calling -&gt; \'\(name)\' and pitch -&gt; \'\(hero)\'
        &quot;&quot;&quot;

        let assistant = colorIdeasModelSession.streamResponse(to: prompt,
                                                            generating: [ColorChip].self,
                                                            includeSchemaInPrompt: true,
                                                            options: creativity.generationOptions)
        
        return assistant
    }
</code></pre>
<p>You can see I check the availability and so there is no doubt here is that code:</p>
<pre class="lang-swift prettyprint-override"><code>    private func checkAvailability(forModel: SystemLanguageModel = .default) throws {
        switch forModel.availability {
        case .available:
            return
        case .unavailable(let reason):
            throw AvailabilityError.unavailable(reason)
        }
    }
</code></pre>
<p>I also have Apple intelligence turned on and my system language settings matching my Siri language.</p>
<p><img src="https://i.sstatic.net/3GQJ1Qxl.png" alt="enter image description here" /></p>
<p>Now, when I google this error Google has its<a href="https://joschua.io/posts/2025/08/23/guardrail-error-xcode-26" rel="nofollow noreferrer"> AI surface this article</a> from someone with a similar issue. I made sure my languages are correct.</p>
<p>To make the issue more infuriating the error does not actually get thrown until I start listening for the stream in my calling code.</p>
<pre class="lang-swift prettyprint-override"><code>    @MainActor
    func generate() async {
        
        self.generatingIdeas = []
        focusField = nil
        
        do {
            
            withAnimation(.snappy) {
                generationStep = .streaming
            }
            
            let ideaStream = try agent.generateIdeas(name: self.projectName, hero: self.projectPitch, creativity: Creativity.allCases.randomElement() ?? .medium, numberOfIdeas: 5)
            
            //Error gets throws from here.
            for try await partialResponse in ideaStream {
                self.generatingIdeas = partialResponse.content
            }
            
            withAnimation(.snappy) {
                generationStep = .finished
            }
            
            generationText = &quot;Generate again&quot;
        }catch {
            generationStep = .error(error)
        }
    }
</code></pre>
<p>What gives? This happens on 2 of my three devices yet I can't find anything on the internet about it? I have <a href="https://feedbackassistant.apple.com/feedback/22632133" rel="nofollow noreferrer">filed it in their Feedback Assistant system</a>.</p>

        
