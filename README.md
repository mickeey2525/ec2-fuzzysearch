# ec2-fuzzysearch

This is a simple Python script that uses fuzzy search to find EC2 instances in your AWS account.
It uses aws cli to get the list of instances and then uses fuzzy search to find the instances you are looking for.

## Installation

```bash
brew tap mickeey2525/ec2-fuzzysearch https://github.com/mickeey2525/ec2-fuzzysearch
brew install ec2-fuzzysearch
```

## Usage
You need to have aws cli installed and configured.
Also, you need to export AWS_REGION and related AWS credentials to use this script.
i.e.

```bash
export AWS_REGION=us-west-2
export AWS_ACCESS_KEY_ID=XXXXXXXXXXXXXXXXXXXX
export AWS_SECRET_ACCESS_KEY=XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
export AWS_SESSION_TOKEN=XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```


```bash
ec2-fuzzysearch
```

Then you can start typing the instance name you are looking for and it will show you the fuzzy search results.
