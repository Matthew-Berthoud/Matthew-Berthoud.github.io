# Matthew-Berthoud.github.io

## Overview
Simple static personal website, whose content is automatically pulled in from my Github.
When a commit is added to main, a github action runs that pulls in information for a few of my github projects.
That is then formatted into HTML templates, re-generating the static content of the site.
That content is then posted as my Github Pages website.

## Setup
```sh
git clone https://github.com/Matthew-Berthoud/personal-vibesite.git
```

### Github Token
If you want to not be quite so rate-limited on your github API requests, you should make authenticated requests.
Create a file called `.env`, and put a [github personal access token](https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/managing-your-personal-access-tokens) for your account in there.
```sh:.env
export GITHUB_TOKEN="YOUR_TOKEN_HERE"

```

