# Contribution

The specification operates on the basis that the Discord API's **responses** are the single source of truth. Contributor-led documentation may provide an idea of how to use the API, but it isn't standardized nor guaranteed to be up-to-date. In contrast, a JSON response from Discord is guaranteed to be consistent with the actual API.

## Submit a JSON

Submit a JSON response from a valid request (REST) or WebSocket (Gateway) event to the respective API version. Only active (non-deprecated) API versions are supported.

### Example

This example uses `curl https://discord.com/api/v9/users/80351110224678912` to submit the JSON response for a `User` resource.

#### Pull Request

```md
## Checklist

- [ ] I have added the JSON to the respective VERSION and FOLDER.
- [ ] I have removed sensitive information _(such as a Discord Bot TOKEN)_ from the provided endpoint below.
- [ ] I have specified whether this pull request refers to an addition, update, or removal.

## Add/Update/Remove JSON

**Endpoint (as specified by the [documentation](https://discord.com/developers/docs/reference)):** `https://discord.com/api/v9/users/{user-id}`

**Discord API Documentation Classifier (for the JSON OBJECT):** `resources/user`
```

_Note that the Discord API Documentation Classifier refers to the User Object; which is also used in other endpoints. Use the classifier to prevent the submission of duplicate JSON objects._

#### Example User Object (Committed)

```json
{
  "id": "80351110224678912",
  "username": "Nelly",
  "discriminator": "1337",
  "avatar": "8342729096ea3675442027381ff50dfe",
  "verified": true,
  "email": "nelly@discord.com",
  "flags": 64,
  "banner": "06c16474723fe537c283b8efa61a30c8",
  "accent_color": 16711680,
  "premium_type": 1,
  "public_flags": 64
}
```

### What is a JSON?

Discord uses JSON to communicate with its clients: [JavaScript Object Notation (JSON)](https://www.json.org/json-en.html) is a data-interchange format. This is the format a bot receives messages in when a request or event occurs.

## Issues

Use issues to identify incomplete features after version 1.

## Discussion

A discussion is used when an issue is not solvable by a pull request.
