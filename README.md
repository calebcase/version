***This tool's maturity is ALPHA. Breaking changes are likely.***

# Version

[![Documentation][godoc.badge]][godoc]
[![Test Status][workflow.tests.badge]][workflow.tests]

This tool automates the process of computing the patch level of the repostory
version. It does this by reading a major, minor version from either a file or
tag and then computing the patch level from the number of commits since the
last commit.

---

[godoc.badge]: https://godoc.org/github.com/calebcase/version?status.svg
[godoc]: https://godoc.org/github.com/calebcase/version
[workflow.tests.badge]: https://github.com/calebcase/version/workflows/tests/badge.svg
[workflow.tests]: https://github.com/calebcase/version/actions?query=workflow%3Atests
