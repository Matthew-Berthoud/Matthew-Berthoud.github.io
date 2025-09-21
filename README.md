# Matthew-Berthoud.github.io

## Overview
Simple static personal website, whose content is automatically pulled in from my Github.
When a commit is added to main, a github action runs that pulls in README information for a few of my github projects.
That is then formatted into HTML templates, re-generating the static content of the site.
That content is then posted as my Github Pages website ([matthewberthoud.com](https://matthewberthoud.com)).

## Setup and Run
To generate an HTML template based on the repo names stored in projects.txt, run the following.
```sh
git clone https://github.com/Matthew-Berthoud/Matthew-Berthoud.github.io.git site
cd site
make
./web
```

### Github Token
If you want to not be quite so rate-limited on your github API requests, you should make authenticated requests.
Create a file called `.env`, and put a [github personal access token](https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/managing-your-personal-access-tokens) for your account in there.
```sh:.env
export GITHUB_TOKEN="YOUR_TOKEN_HERE"

```
For the Github action, this is stored as a Github Secret.

