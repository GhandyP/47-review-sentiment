# 47 — Review sentiment

## Goal

Analyze supplied review text for sentiment, aspects, and evidence while preserving each review ID and avoiding unsupported customer conclusions.

## Run the current prototype

From this case directory:

```bash
export GOOGLE_API_KEY="your-key"
go run .
```

`internal/adkhelper.Run` uses ADK Go v2.2.0, Gemini, and `gemini-flash-latest`; set `ADK_MODEL` to override the model.

Sample prompt:

```text
Analyze these reviews for sentiment and aspects while preserving their IDs.
```

## What exists now

- `main.go (exists)` sends review text to one focused Gemini ADK agent.
- The output is free-form; review ID, sentiment label/score, aspect, evidence span, language, and confidence are not typed.
- This is intentionally text-only: it does not read a review platform, aggregate trends, or respond to customers.

## Missing elements

- **Contract:** Define `ReviewBatch{reviews[{id,text,locale}]}` -> `ReviewAnalysis{id, sentiment, score, aspects[], evidence[], confidence}`.
- **ADK layer:** Add a closed-label schema and batch session; use a future `get_reviews` tool for read-only ingestion, not for replies.
- **Domain boundary:** Add review-platform, language detection, aspect taxonomy, aggregation, and dashboard adapters.
- **Fixtures and validation:** Cover mixed sentiment, sarcasm, multilingual reviews, duplicate IDs, empty text, and aspect overlap; test label consistency.
- **Safety:** Treat results as aggregate signals, strip personal details, and require human review before any customer or employee action.

## Target structure

```text
47-review-sentiment/
├── main.go (exists)
├── README.md (exists)
├── schema.go (next)
├── analyzer.go (next)
├── adapters/reviews.go (next)
├── testdata/reviews.json (next)
└── main_test.go (next)
```

## Implementation order

1. Define review identity, sentiment, aspect, evidence, and confidence schemas.
2. Add locale handling and a batch session with stable IDs.
3. Implement aspect aggregation and duplicate detection.
4. Test mixed sentiment, sarcasm, multilingual text, and missing IDs.
5. Add dashboard/export adapters only as a reviewed, privacy-aware path.

## Done when

- [ ] Every review ID is preserved exactly once with evidence-backed labels.
- [ ] Mixed or ambiguous sentiment is represented rather than forced.
- [ ] Fixtures cover locale, sarcasm, duplicates, and aspect overlap.
- [ ] Individual customer actions never follow automatically from sentiment.

## Next improvement

Make `ReviewAnalysis` ID-preserving and evidence-bearing before adding batch aggregation.
