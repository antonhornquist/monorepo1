# Build graph paths and dependencies

All dependencies of all targets in the workspace including external go dependencies and implicit dependencies can be included in a visualization. The result is an accurate depiction of dependencies but also a very large graph.

```
$ bazel query "deps(//...)" --output graph | dot -Tpng > graph.png && start graph.png
```

![build graph](build-graph-14.png)

All dependencies of all targets in the workspace including external go dependencies but without implicit dependencies can be visualized using the `--noimplicit_deps`. This is still a very large graph.

```
$ bazel query --noimplicit_deps "deps(//...)" --output graph | dot -Tpng > graph.png && start graph.png
```

![build graph](build-graph-15.png)

Narrowing down a use case it's easier to get a usable visualization. Ie. to view all dependencies of the `//eventstore:eventstore` target:

```
$ bazel query "deps(//eventstore:eventstore)" --output graph | dot -Tpng > graph.png && start graph.png
```

![build graph](build-graph-10.png)

This still is a large graph since it includes all required dependencies to build the target. A second argument can be used for `deps()` to limit how many levels of dependencies are included.

Here, four levels:

```
$ bazel query "deps(//eventstore:eventstore, 4)" --output graph | dot -Tpng > graph.png && start graph.png
```

![build graph](build-graph-11.png)

... two levels:

```
$ bazel query "deps(//eventstore:eventstore, 2)" --output graph | dot -Tpng > graph.png && start graph.png
```

![build graph](build-graph-12.png)

Excluding tool dependencies also helps narrow down the visualization if we, for instance, don't care about the Go language toolchain dependency.

```
$ bazel query "deps(//eventstore:eventstore)" --notool_deps --output graph | dot -Tpng > graph.png && start graph.png
```

![build graph](build-graph-13.png)

Excluding all implicit dependencies also simplifies the visualization.

```
$ bazel query "deps(//eventstore:eventstore)" --noimplicit_deps --output graph | dot -Tpng > graph.png && start graph.png
```

![build graph](build-graph-9.png)

Another alternative is to use `filter()` and `except` to exclude targets based on name.

```
$ bazel query "deps(//eventstore:eventstore) except filter(\"@\", deps(//eventstore:eventstore))" --output graph | dot -Tpng > graph.png && start graph.png
```

![build graph](build-graph-4.png)

Using `rdeps()` reverse lookup of target dependants can also be performed, ie. for the `//uniqueid:uniqueid` target:

```
$ bazel query "rdeps(..., //uniqueid:uniqueid)" --output graph | dot -Tpng > graph.png && start graph.png
```

![build graph](build-graph-5.png)

To narrow down a result set Bazel can also be queried for a subset of paths, ie. to visualize all direct and transient dependencies to a particular build target. The below command will visualize the `//eventstore:eventstore` binary target and all its dependencies within the workspace:

```
$ bazel query "allpaths(//eventstore:eventstore, //...)" --output graph | dot -Tpng > graph.png && start graph.png
```

![build graph](build-graph-6.png)

Conversely, the build graph can be queried for all dependants of a particular target. The below command will show all targets in the workspace that directly or indirectly depend on the `//uniqueid:uniqueid` target:

```
$ bazel query "allpaths(//..., //uniqueid:uniqueid)" --output graph | dot -Tpng > graph.png && start graph.png
```

![build graph](build-graph-7.png)

External dependencies are declared up front as part of the build graph and can thus also be analyzed and included in a visualization. The below command will show all targets in the workspace that directly or indirectly depend on the external gin library target:

```
$ bazel query "allpaths(//...,@com_github_gin_gonic_gin//:gin)" --output graph | dot -Tpng > graph.png && start graph.png
```

![build graph](build-graph-8.png)

By querying (ie. using `deps()`, `rdeps()` and `allpaths()`) and filtering (ie. using `filter()`) the build graph can be traversed and analyzed textually or visually depending on a given scenario.

