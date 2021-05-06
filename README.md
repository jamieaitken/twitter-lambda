# twitter-lambda
 This is a small program which is getting used to fetch my profile information from Twitter, specifically the URL of my 
 profile picture. Before this, I had a Lambda written in NodeJS and just updated the code manually. 

The idea was that I could update it on Twitter and then that would update it on [my site](https://jamieaitken.com). I 
also found out that you can have Lambda functions as containers (instead of storing them in an S3 Bucket), so I created 
a Dockerfile and had a mess about with ECR.
