Providers and Publishers that are currently supported
----------------------------------------------------------------------

**Providers(The Platform From where the information is coming from)**

- Github

**Publishers(The Platform where the information will be posted)**

- Twitter
- Github


Before running the bot, you must configure it  so that it can connect to the current supported platforms

Environment variables need for Providers and Publishers
----------------------------------------------------------------------
Github 
----------------------------------------------------------------------
- GITHUB_ACCESS_TOKEN

_NOTE: If you want the content to be published in a README file on a repo, you also need these variables_
- GITHUB_PUBLISH_REPO_OWNER (Your Github username)
- GITHUB_PUBLISH_REPO_NAME (The name of the repo where your README is. It has to be public)
- GITHUB_PUBLISH_REPO_FILE (By default is README)

Twitter
----------------------------------------------------------------------
- TWITTER_CONSUMER_KEY
- TWITTER_CONSUMER_SECRET
- TWITTER_ACCESS_TOKEN
- TWITTER_ACCESS_SECRET

How to setup the environment variables for the platforms
----------------------------------------------------------------------

Github
----------------------------------------------------------------------

To generate the github access tokens follow the given steps

1. Verify your email address, if it hasn't been verified yet.

2. In the upper-right corner of your github profile, click your profile photo, then click Settings.

3. In the left sidebar, click Developer settings.

4. In the left sidebar, click Personal access tokens.

5. Click Generate new token.

6. Give your token a descriptive name.

7. To give your token an expiration, select the Expiration drop-down menu, then click a default or use the calendar picker.

8. Select the scopes, or permissions, you'd like to grant this token. To use your token to access repositories from the command line, select repo.

9. Click Generate token.

For further information click [here](https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/creating-a-personal-access-token)


Twitter
----------------------------------------------------------------------
For getting Twitter keys and secrets click [here](https://developer.twitter.com/en/docs/twitter-api/getting-started/guide) 
