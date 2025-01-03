[![tests](https://github.com/knabben/stalker/actions/workflows/tests.yml/badge.svg)](https://github.com/knabben/stalker/actions/workflows/tests.yml)

## Kubernetes Signal Stalker

Stalk and Hunt for Flake tests on Testgrid Dashboards

Summarizes failures and flakings in the Testgrid board for CI signal enumeration, currently 
fetching `sig-release-master-blocking` and `sig-release-master-informing`

Run the command as `stalker abstract`. A text user interface (TUI) will appear, displaying the combination
of `Board#Tabs` in the first panel. Selecting one of these combinations will show a list of tests in the 
`Tests` section. The two panels below provide the following information:

1. The left panel displays a summary from Slack via the `#release-ci-signal` channel, formatted in Markdown.
2. The right panel shows a GitHub issue, also formatted in Markdown, with the default Kubernetes template pre-filled

To copy for your clipboard the content of the windows pick one of the Windows and press `Ctrl-Space` 
currently only working on WSL2.

![screen](https://github.com/user-attachments/assets/82b55880-dcf5-474c-bd3d-e0f67617a253)

## Features

### AI Augmented Reality

Debugging and identifying the root cause of a problem can be challenging. 
This tool can integrate with well-known LLM machinery APIs and provide automated insights into specific issues.

First set the environment variable `OPENAI_API_KEY` with a token generated in the portal, read the 
[official documentation](https://platform.openai.com/docs/api-reference/authentication) to learn how generate a new
token.

On the bottom panel check the initial prompt and use `Ctrl-B` to call the OpenAI API with the panel text, `Ctrl-Space`
will copy over the final content to your clipboard.

![openai](https://github.com/user-attachments/assets/2f8b3418-1996-443b-810e-992f8ab9ac31)

### GitHub Issue Drafting

It's possible to draft an issue automatically in the [CI Signal Board](https://github.com/orgs/kubernetes/projects/68/views/36).

The Draft issue appears in the DRAFTING section down the first view after the user selecting the panel and pressing `CTRL-b`.

To enable the functionality set a Personal Access Token (PAT) with the proper permissions.
