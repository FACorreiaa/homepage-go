(function () {
  var reduceMotion = window.matchMedia("(prefers-reduced-motion: reduce)").matches;

  function ready(fn) {
    if (document.readyState === "loading") {
      document.addEventListener("DOMContentLoaded", fn);
      return;
    }
    fn();
  }

  ready(function () {
    if (!window.gsap || !window.ScrollTrigger || reduceMotion) {
      document.documentElement.classList.add("motion-ready");
      return;
    }

    gsap.registerPlugin(ScrollTrigger);

    document.querySelectorAll("[data-motion='image-scale']").forEach(function (el) {
      gsap.fromTo(
        el,
        { scale: 0.86, opacity: 0.68 },
        {
          scale: 1,
          opacity: 1,
          ease: "none",
          scrollTrigger: {
            trigger: el,
            start: "top 88%",
            end: "bottom 18%",
            scrub: true
          }
        }
      );
    });

    document.querySelectorAll("[data-motion='pinned-split']").forEach(function (section) {
      var pin = section.querySelector("[data-motion-pin]");
      var track = section.querySelector("[data-motion-track]");
      if (!pin || !track || window.innerWidth < 1024) {
        return;
      }

      ScrollTrigger.create({
        trigger: section,
        start: "top top",
        end: "bottom bottom",
        pin: pin,
        pinSpacing: false
      });

      gsap.fromTo(
        track,
        { yPercent: 4 },
        {
          yPercent: -4,
          ease: "none",
          scrollTrigger: {
            trigger: section,
            start: "top bottom",
            end: "bottom top",
            scrub: true
          }
        }
      );
    });

    document.querySelectorAll("[data-motion='stack']").forEach(function (stack) {
      var cards = gsap.utils.toArray(stack.querySelectorAll("[data-stack-card]"));
      cards.forEach(function (card, index) {
        gsap.to(card, {
          y: -18 * index,
          scale: 1 - index * 0.018,
          ease: "none",
          scrollTrigger: {
            trigger: card,
            start: "top 74%",
            end: "top 24%",
            scrub: true
          }
        });
      });
    });

    document.querySelectorAll("[data-motion='text-reveal']").forEach(function (el) {
      var words = el.textContent.trim().split(/\s+/);
      el.textContent = "";
      words.forEach(function (word, index) {
        var span = document.createElement("span");
        span.textContent = word + (index === words.length - 1 ? "" : " ");
        span.style.opacity = "0.16";
        el.appendChild(span);
      });

      gsap.to(el.querySelectorAll("span"), {
        opacity: 1,
        stagger: 0.04,
        ease: "none",
        scrollTrigger: {
          trigger: el,
          start: "top 82%",
          end: "bottom 36%",
          scrub: true
        }
      });
    });

    document.documentElement.classList.add("motion-ready");
  });
})();
