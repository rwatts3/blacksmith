---
title: F.A.Q.
enterprise: false
---

# F.A.Q.

## Can I use Blacksmith with an existing data pipeline?

Yes. Blacksmith is flexible and allows you to have external dependencies or pieces
of software in addition to the application built on top of it.

We are well aware analytics and marketing teams already use third-party services
like Segment and Zapier. Blacksmith can act as an addition or a substitute to these
services depending on your needs.

## What are the prerequisites to learn Blacksmith?

Blacksmith is developed with the Go language. Therefore it is necessary to have
some experience with it.

The Blacksmith interfaces make it very easy to dive quickly into development.
Even if you just experienced Go for a few days it should be enough to understand
how Blacksmith works and create a first simple data pipeline.

## Why did you choose Go?

Go has the right level of simplicity and abstraction we desired to build such a
product. Its performances and design choices lead us to pick it from the start,
without any regret.

Also, Go has become "the language of the cloud" in the past years. A major part
of cloud infrastructures and tools rely on Go such as Docker, Kubernetes, and
Terraform.

## What is an adapter?

An adapter is an *implementation* of an *interface*. For example the PostgreSQL
store adapter let you use PostgreSQL as a store for Blacksmith.

## What is the license?

The Go public APIs are [available on GitHub](https://github.com/nunchistudio/blacksmith),
and licensed under the [Apache License 2.0](https://github.com/nunchistudio/blacksmith/blob/master/LICENSE).

The use of Blacksmith is governed by the
[Blacksmith Terms and Conditions](http://nunchi.studio/blacksmith/legal/terms).