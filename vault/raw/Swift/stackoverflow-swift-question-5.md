---
source: Stack Overflow RSS Feed
ingested_at: 2026-05-11T08:06:16.067179
type: web
status: uncompiled
---

# Failed to load interstitial ad with error: Request Error: No ad to show

**Stack Overflow Question**

**Link:** https://stackoverflow.com/questions/79933459/failed-to-load-interstitial-ad-with-error-request-error-no-ad-to-show

**Published:** 2026-04-29T06:40:12Z

**Updated:** 2026-05-11T02:07:05Z

**Author:** user3318703

**Tags:** ios, swift, sdk, admob, firebase-admob

**Summary:**

            <p><strong>Environment</strong><br />
<strong>SDK Version</strong>: 13.2.0 (via SPM)[https://github.com/googleads/swift-package-manager-google-mobile-ads.git]</p>
<p><strong>Xcode Version</strong>: 26.0</p>
<p><strong>Platform</strong>: iOS</p>
<p><strong>Ad Format</strong>: Interstitial</p>
<p><strong>Issue Description</strong><br />
I am encountering a persistent &quot;Request Error: No ad to show&quot; when running my app in a production environment. The implementation works perfectly in Debug mode (using test ad unit IDs), where ads load and display as expected. However, once switched to production Ad Unit IDs and deployed, the load request fails consistently.</p>
<p><strong>Full error log</strong></p>
<p>Error Domain=com.google.admob Code=1 &quot;Request Error: No ad to show.&quot; UserInfo={NSLocalizedDescription=Request Error: No ad to show.</p>
<p><strong>Implementation Code</strong></p>
<pre><code>InterstitialAd.load(
    with: &quot;&lt;INTERSTITIAL_ADS_UNIT_ID&gt;&quot;,
    request: request,
    completionHandler: { [weak self] ad, error in
        guard let self = self else { return }
                
        if let error = error {
            print(&quot;[AdManager] Failed to load interstitial ad with error: \(error.localizedDescription)&quot;)
            return
        }
                
        // Ensure UI updates happen on main actor
        DispatchQueue.main.async {
            self.interstitial = ad
            self.interstitial?.fullScreenContentDelegate = self
            print(&quot;[AdManager] Interstitial ad loaded successfully&quot;)
        }
    }
)
</code></pre>
<p><strong>Troubleshooting Steps Taken</strong></p>
<p>App ID/Ad Unit ID Verification: Confirmed that the IDs match the AdMob dashboard exactly.​</p>
<p>Test Devices: Verified that the behavior persists across different physical devices.</p>
<p>App Store Status: The app is linked to the AdMob dashboard and the app-ads.txt is verified<code> ​and &quot;Authorized.&quot;</code></p>
<p>Payments/Policy: No policy violations or payment holds are present in the AdMob console.</p>
<p>​Thanks in advance</p>

        
