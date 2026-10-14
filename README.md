# TeamCity Go Test Runner

`teamcity-go-test` is a replacement test runner for TeamCity for Go. It is
heavily inspired by [Pavel Gulbin's work][1] <other repo>, but modified to fit
our use case, where we have lots of tests which take a long time but are safe
to run in parallel. The workflow is:

1. Compile a test binary using `go test -c`
2. Pipe a list of test names, one per line, into `teamcity-go-test` , with the
   `-test` parameter pointing to the executable. To run tests in parallel, set
   `-parallelism` to a number greater than 1.

## TeamCity Parallel Tests

Pass `-parallel-tests` when the build uses TeamCity's [Parallel Tests][2]
feature. The runner then:

- reads the `teamcity.build.parallelTests.excludesFile` system property from
  the file named by `TEAMCITY_BUILD_PROPERTIES_FILE`, and skips the tests
  TeamCity assigned to other batches. If the exclusions file cannot be read
  the runner fails rather than running every test. Without the property (the
  build is not batched) every test runs.
- reports each test as `TestName.run`. TeamCity splits batches by test class
  (the text before the last `.`), so this makes each Go test function its own
  class while the exclusions still contain plain Go test names.


[1]: https://github.com/2tvenom/go-test-teamcity
[2]: https://www.jetbrains.com/help/teamcity/parallel-tests.html
