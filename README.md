# Advent Of Code Template

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Language agnostic template repository for Advent Of Code challenges, providing basic automation.

## Automations

### Dependabot

This repository is setup to enable Dependabot PRs for all your project dependencies. The following update configurations are provided;

1. Github Actions updates - This will automatically keep your action versions up to date as new major versions are released
2. Docker updates - This will automatically keep any docker files, such as VSCode Devcontainers, up to date.
3. Your own package ecosystem - This is set to `npm` by default, but you can find all the possible values to change it to [in the Dependabot documentation](https://docs.github.com/en/code-security/dependabot/dependabot-version-updates/configuration-options-for-the-dependabot.yml-file#package-ecosystem)

### Automerge

The [automerge](./.github/workflows/automerge.yml) workflow runs on every Dependabot PR and will automatically merge it if your required checks pass. In its initial form, this repository **has no default checks**, thus these PRs will always be merged.

### Collect Input

This workflow will automatically retrieve and commit the current day's puzzle input from the [Advent of Code](https://adventofcode.com) website.

For this to work, you must first add your `AOC_SESSION_COOKIE` to the repository secrets. You can find this as so;

1. Visit https://adventofcode.com
2. If you are not already logged in, log in with your account of choice
3. Open your browser's Developer Tools (usually by pressing F12 on your keyboard)
4. Navigate to the cookies stored on this site (in Chrome, it's under `Application`)
5. Find the cookie named `session` and copy its value

This workflow will run on every day of the challenge, from 1st-25th December with a bonus day on [Boxing Day](https://en.wikipedia.org/wiki/Boxing_Day) (26th December). On the 26th, the workflow will automatically update itself to turn off daily input collection, otherwise it will attempt to find inputs next year and fail.
