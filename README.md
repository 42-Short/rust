# rust
Rust Piscine built using [Nils Mathieu's subjects](https://github.com/nils-mathieu/piscine-rust) and the [shortinette framework](https://pkg.go.dev/github.com/42-Short/shortinette).

## Usage
### Environment Configuration
Required variables:
**Github API**:
* `GITHUB_ADMIN`: GitHub username of the account which is to be used for creating repos and commiting.
* `GITHUB_EMAIL`: GitHub email of said account.
* `GITHUB_ORGANISATION`: Name of the GitHub organisation under which the repositories are to be created.
* `GITHUB_TOKEN`: GitHub personal access token with admin permissions on the organisation.
**Webhook**:
* `WEBHOOK_URL`: <IP>:<PORT>/<ENDPOINT> (e.g., `12.34.56.78:1234/webhook`). Note that for GitHub to be able to send webhook payloads, you will need to have a **public IP address**.
**Participants Configuration**:
* `CONFIG_PATH`: Path to your `.json` file containing the participants information (see [below](#participants-configuration) for details).
### Participants Configuration
This configuration is in `json` format and contains a list of participants. Each participant has two fields:
* `github_username`: GitHub user, will be given write permissions to the created repo.
* `intra_login`: Unique identifier, will be used for naming the repos (`repo_name = fmt.Sprintf("%s-%s", participant.IntraLogin, module.Name)).
```json
{
    "participants": [
        {
            "github_username": "shortinette-test",
            "intra_login": "1"
        }
    ]
}
```
### Run That Shit
```zsh
go run .
```

