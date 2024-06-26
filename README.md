# ec2-fuzzysearch

This is a simple CLI that searches for EC2 instances in your AWS account using the Name Tag.
It utilizes the AWS SDK for Go to retrieve the list of instances that has Name tags, then applies fuzzy search to help you find the instance you are looking for.

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

or you can use aws profile for sso users.
i.e.
```bash
export AWS_PROFILE=sso
```

Then you can run the following command to search for the instances.

```bash
ec2-fuzzysearch -p <profile> -r <region>
```

Then you can start typing the instance name you are looking for and it will show you the fuzzy search results.
