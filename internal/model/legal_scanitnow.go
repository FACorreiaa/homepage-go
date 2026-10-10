package model

// scanitnowLegal: ScanItNow, the iPhone shopping-total app (repo FACorreiaa/scanit). v1 has no account,
// no analytics SDK and no purchases, so the App Store privacy label is "Data Not Collected"; keep this
// policy in step with that label. Update both when v1.1 adds PostHog/Sentry or the one-time unlock.
func scanitnowLegal() LegalApp {
	const (
		operator = "Fernando Correia (FC Software Studio)"
		contact  = "fernando@facorreia.com"
		updated  = "10 October 2026"
	)
	app := LegalApp{
		Slug:          "scanitnow",
		Name:          "ScanItNow",
		Operator:      operator,
		SupportEmail:  contact,
		PrivacyEmail:  contact,
		WebsiteURL:    "https://facorreia.com/projects/scanitnow",
		SupportIntro:  "Email is the fastest route. Tell us your iPhone model and iOS version, and, if a product scanned wrongly, the barcode number printed under it.",
		DeleteHeading: "Delete your data",
		DeleteBody:    "ScanItNow has no account: your trips, items and prices live only on your iPhone. Deleting the app deletes all of it.",
	}
	app.Privacy = LegalDocument{
		App: "scanitnow", AppName: "ScanItNow", Kind: "privacy",
		Title:       "ScanItNow Privacy Policy",
		LastUpdated: updated,
		Intro:       "ScanItNow adds up your shopping as you go. It is operated by " + operator + ". It works without an account, and we do not collect personal data.",
		Sections: []LegalSection{
			{Heading: "What stays on your iPhone", Body: []string{
				"Your trips, the items in them, the prices you confirm, your spending limits and the stores you name are stored only on your iPhone. We have no copy and no server that receives them.",
				"The camera reads barcodes on the device. No photo or video is saved or sent anywhere.",
			}},
			{Heading: "Product lookups", Body: []string{
				"When you add a product, its barcode number is sent to Open Food Facts (openfoodfacts.org), a free public product database, to fetch the product's name and photo. Only the barcode number is sent: never the price, the store, your location or anything that identifies you. Open Food Facts' own privacy policy applies to that request.",
			}},
			{Heading: "What we do not do", Body: []string{
				"ScanItNow has no analytics or advertising SDKs, does not track you across apps or websites, does not use your location, and does not sell or share data.",
				"If you share crash reports and app analytics with developers in your iPhone's settings, Apple provides us anonymous, aggregated crash and usage statistics. That is controlled by you in Settings > Privacy & Security > Analytics & Improvements.",
			}},
			{Heading: "Children", Body: []string{
				"ScanItNow is not directed at children under 13 and does not knowingly process their personal data.",
			}},
			{Heading: "Deleting your data", Body: []string{
				"Deleting the app removes everything it stored. There is nothing to delete on our side.",
			}},
			{Heading: "Your rights and contact", Body: []string{
				"Under the GDPR you may contact us about any privacy question at " + contact + ". You can also complain to your national data-protection authority (in Portugal, the CNPD).",
				"If this policy changes, for example when a future version adds an optional account or analytics, the date above changes and the App Store privacy label is updated before that version ships.",
			}},
		},
	}
	app.Terms = LegalDocument{
		App: "scanitnow", AppName: "ScanItNow", Kind: "terms",
		Title:       "ScanItNow Terms of Use",
		LastUpdated: updated,
		Intro:       "These Terms govern your use of ScanItNow (the \"App\") on iPhone, operated by " + operator + ". By downloading or using the App you agree to them.",
		Sections: []LegalSection{
			{Heading: "Licence", Body: []string{
				"We grant you a personal, non-transferable, revocable licence to use the App on Apple devices you own or control, in accordance with the Apple Media Services Terms and these Terms.",
			}},
			{Heading: "Totals are estimates", Body: []string{
				"ScanItNow adds up the prices you scan or type. Its totals are only as accurate as those prices. The amount charged at the till is what you pay.",
			}},
			{Heading: "Product information", Body: []string{
				"Product names and photos come from Open Food Facts, a community database. They may be missing or wrong, and we do not guarantee them.",
			}},
			{Heading: "Acceptable use", Body: []string{
				"Use the App for your own shopping, lawfully and in line with each store's rules. Do not attempt to disrupt or reverse-engineer it.",
			}},
			{Heading: "Disclaimer and liability", Body: []string{
				"The App is provided \"as is\" without warranties of any kind. To the maximum extent permitted by law, " + operator + " is not liable for indirect or consequential damages arising from relying on its totals, prices or product information. Nothing limits liability that cannot be excluded by law.",
			}},
			{Heading: "Changes", Body: []string{
				"We may update these Terms; the date above changes when we do, and continued use after a change means you accept it.",
			}},
			{Heading: "Governing law", Body: []string{
				"These Terms are governed by the laws of Portugal, without regard to conflict-of-laws rules. Consumers in the EU keep the protections of their country of residence.",
			}},
			{Heading: "Contact", Body: []string{"Questions about these Terms: " + contact + "."}},
		},
	}
	return app
}
