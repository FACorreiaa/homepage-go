---
title: "Uploads, a 3D activity view, and Loci on Telegram"
summary: "Norviq now takes a CSV or a screenshot instead of a form. Khepri draws activities in 3D. Loci answers on Telegram and ships pictures with the places."
category: engineering
date: "2026-09-11"
---

Three things landed today, one per product. The through-line is the same in all of them: stop asking people to type what they already have.

## Norviq: upload the file, skip the form

Adding an expense or updating a portfolio in Norviq meant a form. Ticker, quantity, price, date, repeat. Nobody does that for twenty positions, so the data went stale and the app got less useful every week.

Two import paths now exist instead.

The first is CSV. Broker exports go straight in through the import service, get parsed into positions, and land as real rows. This is the boring path and the one that scales past a handful of holdings.

The second is a screenshot. You take a picture of your broker app or a receipt and upload it. A vision model reads it, an extractor turns the text into structured positions or expense lines, and you confirm before anything is written. Receipts go through the same shape: upload, OCR, structured line items, confirm.

I am deliberately keeping the confirm step. Vision extraction is good, not perfect, and silently writing a wrong cost basis into a portfolio is worse than one extra tap. There is an eval script in the repo for exactly this reason — I want a number on extraction quality, not a feeling.

Both paths share the same wire format into the AI layer, so there is one code path for "here is an image, give me structured data back" rather than one per feature.

## Khepri: activities in three dimensions

Khepri has had activity data for a while — workouts, sessions, whatever Strava sends over. It was rendered as charts, which is fine for one metric over time and useless for seeing the shape of a month.

There is now a 3D view of activities. Volume, intensity, and time on three axes instead of stacking flat charts. It makes gaps obvious in a way a line chart hides: you can see the two weeks you did nothing without reading a single number.

This is the kind of thing that sounds like decoration and is not. The point of a coach that remembers is that the record tells you something. If the record is only legible to a chart library, it is not telling you much.

## Loci: Telegram, and places that have pictures

Loci had a web client and an MCP server. It now has a Telegram bot too.

The bot registers by webhook and dispatches straight into the same itinerary pipeline the web uses. You message it a city and a mood, it comes back with an itinerary. No separate planner, no second set of prompts to keep in sync — one service, three surfaces.

It did not work on the first message. The bot wedged on its first inbound update and answered nowhere useful, which took a fix to get it replying in the chat it was messaged from. Webhook handlers are always more state than they look.

The other half is pictures. POI citations used to point at rows that did not always exist, and places came back as text. Citations now resolve to real POI rows, and those rows carry images — so an itinerary arrives with photos of the places in it, on Telegram and on the web. Generating and attaching the imagery happens in the chat service, next to where the POIs are resolved, so both surfaces get it for free.

## What is not done

The Norviq extraction quality number does not exist yet — the eval script is written, the run is not. Khepri's 3D view is one visualization, not a dashboard. Loci's bot does itineraries; it does not do the rest of the product.

All three are on the same principle though: the input people actually have is a file, a screenshot, or a message — not a form.
