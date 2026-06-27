---
source: Stack Overflow RSS Feed
ingested_at: 2026-05-11T08:06:16.066394
type: web
status: uncompiled
---

# How to add shadow to border?

**Stack Overflow Question**

**Link:** https://stackoverflow.com/questions/79939205/how-to-add-shadow-to-border

**Published:** 2026-05-11T07:29:30Z

**Updated:** 2026-05-11T08:02:46Z

**Author:** steveSarsawa

**Tags:** ios, swift, uibutton

**Summary:**

            <p>I'm trying to add shadow to Border, but it is not working. Shadow should as seen in image, but actual requirement is like Button will be clear background, and shadow around border. So how can I achieve it?</p>
<p><a href="https://i.sstatic.net/MBGuDULp.png" rel="nofollow noreferrer"><img src="https://i.sstatic.net/MBGuDULp.png" alt="" /></a></p>
<pre><code>
  class GlowButton: UIButton {  
        
      private let borderGlowLayer = CAShapeLayer()  
        
      override init(frame: CGRect) {  
          super.init(frame: frame)  
          setup()  
      }  
        
      required init?(coder: NSCoder) {  
          super.init(coder: coder)  
          setup()  
      }  
        
      private func setup() {  
          backgroundColor = .clear  
          layer.masksToBounds = false  
          clipsToBounds = false  
            
          borderGlowLayer.fillColor   = UIColor.clear.cgColor  
          borderGlowLayer.strokeColor = UIColor.clear.cgColor  
          borderGlowLayer.lineWidth   = 2.5  
          borderGlowLayer.shadowOffset  = .zero  
          borderGlowLayer.shadowRadius  = 0  
          borderGlowLayer.shadowOpacity = 0  
          borderGlowLayer.masksToBounds = false  
          layer.addSublayer(borderGlowLayer)  
      }  
        
      override func layoutSubviews() {  
          super.layoutSubviews()  
          layer.masksToBounds = false  
          clipsToBounds = false  
          borderGlowLayer.frame = bounds  
          borderGlowLayer.path  = UIBezierPath(roundedRect: bounds, cornerRadius: 20).cgPath  
      }  
        
      override func didUpdateFocus(in context: UIFocusUpdateContext, with coordinator: UIFocusAnimationCoordinator) {  
          super.didUpdateFocus(in: context, with: coordinator)  
          coordinator.addCoordinatedAnimations {  
              if self.isFocused {  
                  self.borderGlowLayer.strokeColor  = Colors.focused_border.cgColor  
                  self.borderGlowLayer.shadowColor  = Colors.focused_border.cgColor  
                  self.borderGlowLayer.shadowRadius  = 12  
                  self.borderGlowLayer.shadowOpacity = 1.0  
              } else {  
                  self.borderGlowLayer.strokeColor   = UIColor.clear.cgColor  
                  self.borderGlowLayer.shadowOpacity = 0.0  
                  self.borderGlowLayer.shadowRadius  = 0  
              }  
          }  
      }  
  }
</code></pre>

        
