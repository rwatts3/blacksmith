---
title: How it works
enterprise: false
---

# How it works

Wether it is your first time doing data engineering or not, this guide is important
since it describes some specifities of how Blacksmith works.

Let's say that, for now, Blacksmith should look like this to you: a blackbox.

![Step 01](/images/blacksmith/how.001.png)

Some things happen, but you don't know why or how.

## Sources and destinations

The role of a data pipeline is mainly to take care of Extracting data from *sources*,
Transforming it when necessary, and Loading it to *destinations*. In short: doing ETL.
Each source has *triggers* and each destination has *actions*.

In other words, the blackbox that we imagined earlier is able to receive events from
sources' triggers, transform data of these events, and run it against the appropriate
destinations' actions.

Sources can be websites, mobile applications, third-party services, databases, etc.
Destinations can be databases, data warehouses, or third-party services for analytics
or marketing.

![Step 02](/images/blacksmith/how.002.png)

Triggers and actions inherit properties of their parent source or destination — such
as retry logic — but can override it in special cases.

Both triggers and actions have a business logic for handling enrichment, transformation,
and automating the data flow.

## The gateway and scheduler

Instead of using a lot of micro-services and over-engineer the whole process, Blacksmith
keep your data pipeline simple with only two services: the *gateway* and the *scheduler*.

The gateway is in charge of extracting the events from sources on triggers. It can
happen on:
- HTTP requests;
- CRON schedules;
- CDC notifications;
- Pub / Sub messages.

The scheduler is in charge of loading the events into actions of destinations.

![Step 03](/images/blacksmith/how.003.png)

These services can run on a single machine but should be splitted for production
use. This way, you gain more flexibility, security, and scalability.

The services can be `standard` or `enterprise`, depending on the Blacksmith Edition
you rely on.

## Database-as-a-Queue

The gateway needs a way to keep track of received events and *jobs* to execute onto
destinations. The scheduler needs a way to keep track of job status (also called
*transitions*) so it can be aware of successes, failures, and discards.

The best way to achieve this is to have a persistent *store*. Every events received
by the gateway are stored along the desired jobs. The scheduler knows how and when
to load the data to appropriate destinations and can keep a record of jobs'
transitions. 

![Step 04](/images/blacksmith/how.004.png)

The `store` adapter is the only one required along the services.

Available `store` adapters:
- PostgreSQL (`postgres`)

## Enabling realtime

The store is perfect to persist events, jobs, and status. But we often need to
send data from sources to destinations in realtime.

We accomplish this by adding a pub / sub mechanism between the gateway and scheduler.
Once the event is received and stored, it is then automatically *published*  by the
gateway to the scheduler that *subscribed* to the events.

![Step 05](/images/blacksmith/how.005.png)

The `pubsub` adapter is optional. When no adapter is provided or the action does not
handle realtime events, the scheduler will send data schedule retries given (order
matters):
- the destination's action schedule; or
- the destination's schedule; or
- the default schedule.

Once configured, the `pubsub` adapter allows realtine message extraction from
queues / topics / subscriptions.

Available `pubsub` adapters:
- Azure Service Bus (`azure`)
- Google Pub / Sub (`google`)
- AWS SNS / SQS (`aws`)
- Kafka (`kafka`)
- NATS (`nats`)
- RabbitMQ (`rabbitmq`)

## Versioning migrations

Data pipelines often — to not say always — need to communicate with databases.
Managing and versioning database migrations is a difficult task to achieve,
especially across distributed teams in an organization. Also, the data schemas
of the data pipeline and the one of the databases need to be synchronized to
avoid schema incompatibilities as much as possible.

To resolve this challenge we add a *wanderer*. The wanderer does not handle the
migrations logic but keeps track of migration runs across every adapters, sources,
triggers, destinations, and actions. It is therefore aware of what migrations need
to run and which ones to rollback.

For example, a destination can have migrations and each of its actions can also
have specific migrations. Sources and destinations can have a `Migrate` function
which defines the migration logic to run for the adapter and will be executed
when running a migration against the adapter.

Migrations have a *up* and *down* logic allowing rollbacks. The wanderer leverages
the *supervisor* as described below so only one migration can run at a time on a
given target.

The `wanderer` adapter is optional. It is only needed when managing migrations from
a Blacksmith application.

Available `wanderer` adapters:
- PostgreSQL (`postgres`)

> The *wanderer* is only available using Blacksmith Enterprise Edition.

## Distributed environments

Most companies need to have zero downtime to maximize their service availability.

Blacksmith applications can optionally leverage a *supervisor*, which is in charge
of acquiring and releasing lock accesses in distributed environments. This ensures
resources in a multi-nodes cluster are accessed by a single instance of the
gateway, scheduler, and CLI to avoid collision when listening for events, executing
jobs, or running migrations.

The `supervisor` adapter is optional. It is only needed when running Blacksmith
applications in distributed environments.

Available `supervisor` adapters:
- Consul (`consul`)

> The *supervisor* is only available using Blacksmith Enterprise Edition.

## Conclusion

By splitting Blacksmith into two services only, we ensure simplicity and consistency.
This also add a great separation of concerns for better security and scalability.
While the gateway may be exposed to the outside world, the scheduler can live
inside a private network with no public address.

The store keeps track of everything happening in the data pipeline. The pub / sub
allows realtime events forwarding between the gateway and scheduler.

The wanderer keeps track of migrations across every sources and destinations.

The supervisor allows to run Blacksmith applications in distributed environments.