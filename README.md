# openai-cli

🚧 hobby project

cf. [OpenAI API reference](https://platform.openai.com/docs/api-reference)

## Models

- [x] GET /v1/models
- [x] GET /v1/models/{model}

## Completions

- [x] POST /v1/completions
- [x] POST /v1/chat/completions

## Edits

- [x] POST /v1/edits

## Images

- [x] POST /v1/images/generatoins
- [ ] POST /v1/images/edits
- [ ] POST /v1/images/variations

## Embeddings

- [ ] POST /v1/embeddings

## Audio

- [ ] POST /v1/audio/transcriptions
- [ ] POST /v1/audio/translations

## Files

- [x] GET /v1/files
- [ ] GET /v1/files/{file_id}
- [ ] GET /v1/files/{file_id}/content
- [x] POST /v1/files
- [ ] DELETE /v1/files/{file_id}

## Fine-tunes

- [ ] GET /v1/fine-tunes
- [ ] GET /v1/fine-tunes/{fine_tune_id}
- [ ] GET /v1/fine-tunes/{fine_tune_id}/events
- [ ] POST /v1/fine-tunes
- [ ] POST /v1/fine-tunes/{fine_tune_id}/cancel
- [ ] DELETE /v1/models/{model}

## Moderations

- [x] POST /v1/moderations
