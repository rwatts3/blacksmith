# Blacksmith

Blacksmith is a low-code ecosystem offering a complete and consistent approach for
building self-managed data engineering solutions. It allows organizations to process,
trust, and visualize all their data flowing from end-to-end in a consistent way.

Any team that is building — or think about building — such a platform knows the
tremendous amount of work needed to properly accomplish this mission. Think of
Blacksmith as the central piece of your data engineering workflow, leading you to
save months of customized and professional work.
 
![Data engineering with Blacksmith](https://nunchi.studio/images/blacksmith/approach.png)

By leveraging Blacksmith, organizations benefit a single source of truth for all
their data with a unique developer experience.

## Production-ready integrations

![Blacksmith integrations](https://nunchi.studio/images/blacksmith/integrations.png)

## Product offerings

**Blacksmith is not an open-source software.** This repository only holds the
public Go APIs, allowing organizations to build reliable data engineering solutions
on top of Blacksmith using Go. Blacksmith itself is built and distributed as a CLI
and as [a Docker image](https://github.com/nunchistudio/blacksmith-docker).

[Blacksmith is available in two Editions](https://nunchi.studio/blacksmith/pricing):
- **Blacksmith Standard Edition** addresses the technical complexity of data
  engineering. It is and will always be free. This Edition focuses on building
  ETL platforms for better data quality and stronger data reliability across systems.
- **Blacksmith Enterprise Edition** addresses the complexity of collaboration and
  governance across multi-team and multi-scope data solutions. This Edition adds
  advanced features for the enterprises, such as migrations management, distributed
  semaphore, and a built-in dashboard.

## More than just ETL

The **Enterprise Edition** offers advanced features for more complete data
engineering solutions.

**Granular migrations management:**
```bash
$ blacksmith migrations rollback --version 20200930071321 \
  --scope "destination:sqlike(crm)" \
  --scope "destination:sqlike(warehouse)"
```

**Powerful REST API:**
```bash
$ curl --request GET --url 'https://example.com/admin/api/store/jobs' \
  -d events.sources_in=cms \
  -d events.sources_in=crm \
  -d jobs.destinations_in=warehouse \
  -d jobs.actions_in=register \
  -d jobs.status_in=discarded \
  -d offset=0 -d limit=100
```

**Built-in dashboard:**
![Blacksmith Dashboard](https://nunchi.studio/images/blacksmith/dashboard.002.png)

## Professional services

Along consulting and training, we provide different product offerings as well as
different levels of support.

- [Request a demo](https://nunchi.studio/blacksmith/forms/demo)
- [Discover our services](https://nunchi.studio/support)

## License

Repository licensed under the [Apache License, Version 2.0](./LICENSE).

By downloading, installing, and using Blacksmith, you agree to the
[Blacksmith Terms and Conditions](https://nunchi.studio/legal/terms).
