---
kep-number: draft-20190214
title: Dynamic CRDs
authors:
  - "@gerred"
owners:
  - TBD
editor: TBD
creation-date: 2019-02-14
last-updated: 2019-02-14
status: provisional
---

# Framework-specific Dynamic Custom Resource Definitions

## Table of Contents

- [Framework-specific Dynamic Custom Resource Definitions](#framework-specific-dynamic-custom-resource-definitions)
  - [Table of Contents](#table-of-contents)
  - [Summary](#summary)
  - [Motivation](#motivation)
    - [Goals](#goals)
    - [Non-Goals](#non-goals)
  - [Proposal](#proposal)
    - [Risks and Mitigations](#risks-and-mitigations)
  - [Graduation Criteria](#graduation-criteria)
  - [Implementation History](#implementation-history)

## Summary

This KEP aims to make the end user experience for frameworks more specific to the business domain they represent. By implementing dynamic CRDs, frameworks will be able to represent their components in a declarative way, specific to their framework. Additionally, it enables framework developers to provide day 2 operations tasks as CRDs, complete with their own plans and tasks.

## Motivation

Currently, the interface for using frameworks in KUDO is very generic. Users create an `Instance` CRD with specs. Operator developers and users expect to be able to use contextual business objects for their operators instead of generic objects. This enables a more focused experience for users of KEP.

The goal of this KEP is to improve the end user UX through dynamic CRDs. Other than the ability to specfiy CRDs, and adjusting existing framework development CRDs to accomodate this change, it is not the goal of this KEP to change the framework development UX.

### Goals

- Create a mechanism for framework developers to specify a CRD
- Enable management for custom resources based on dynamic CRDs. Deploying a framework specific custom resource should deploy a plan as `Instance` was able to before.

### Non-Goals

- Change the framework developer UX for templates, parameters, tasks, and plans.

## Proposal

This is where we get down to the nitty gritty of what the proposal actually is.

### Risks and Mitigations

What are the risks of this proposal and how do we mitigate.
Think broadly.
For example, consider both security and how this will impact the larger kubernetes ecosystem.

## Graduation Criteria

How will we know that this has succeeded?
Gathering user feedback is crucial for building high quality experiences and owners have the important responsibility of setting milestones for stability and completeness.
Hopefully the content previously contained in [umbrella issues][] will be tracked in the `Graduation Criteria` section.

[umbrella issues]: https://github.com/kubernetes/kubernetes/issues/42752

## Implementation History

- `Summary`, `Motivation`, and `Goals` being merged.
