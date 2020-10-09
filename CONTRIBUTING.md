# Contribute Code

You are welcome to contribute to project PaddlePaddle EDL. To contribute to EDL, you have to agree with the 
[PaddlePaddle EDL Contributor License Agreement](https://gist.github.com/guru4elephant/d0207d374c0ddfa11b70751dca97703b).

We sincerely appreciate your contribution.  This document explains our workflow and work style.

## Workflow

PaddlePaddle EDL uses this [Git branching model](http://nvie.com/posts/a-successful-git-branching-model/).  The following steps guide usual contributions.

1. Fork

   Our development community has been growing fastly; it doesn't make sense for everyone to write into the official repo.  So, please file Pull Requests from your fork.  To make a fork,  just head over to the GitHub page and click the ["Fork" button](https://help.github.com/articles/fork-a-repo/).

1. Clone

   To make a copy of your fork to your local computers, please run

   ```bash
   git clone https://github.com/your-github-account/edl
   cd paddle
   ```

1. Create the local feature branch

   For daily works like adding a new feature or fixing a bug, please open your feature branch before coding:

   ```bash
   git checkout -b my-cool-stuff
   ```

1. Commit

   Before issuing your first `git commit` command, please install [`pre-commit`](http://pre-commit.com/) by running the following commands:

   ```bash
   pip install pre-commit
   pre-commit install
   ```

   Our pre-commit configuration requires clang-format 3.8 for auto-formating C/C++ code and yapf for Python.

   Once installed, `pre-commit` checks the style of code and documentation in every commit.  We will see something like the following when you run `git commit`:

   ```
   ➜  git commit
   CRLF end-lines remover...............................(no files to check)Skipped
   yapf.................................................(no files to check)Skipped
   Check for added large files..............................................Passed
   Check for merge conflicts................................................Passed
   Check for broken symlinks................................................Passed
   Detect Private Key...................................(no files to check)Skipped
   Fix End of Files.....................................(no files to check)Skipped
   clang-formater.......................................(no files to check)Skipped
   [my-cool-stuff c703c041] add test file
    1 file changed, 0 insertions(+), 0 deletions(-)
    create mode 100644 233
   ```

	NOTE: The `yapf` installed by `pip install pre-commit` and `conda install -c conda-forge pre-commit` is slightly different. Paddle developers use `pip install pre-commit`.

1. Build and test

   Users can build EDL natively on Linux.  But to unify the building environment and to make it easy for debugging.

1. Keep pulling

   An experienced Git user pulls from the official repo often -- daily or even hourly, so they notice conflicts with others work early, and it's easier to resolve smaller conflicts.

   ```bash
   git remote add upstream https://github.com/PaddlePaddle/edl
   git pull upstream develop
   ```

1. Push and file a pull request

   You can "push" your local work into your forked repo:

   ```bash
   git push origin my-cool-stuff
   ```

   The push allows you to create a pull request, requesting owners of this [official repo](https://github.com/PaddlePaddle/edl) to pull your change into the official one.

   To create a pull request, please follow [these steps](https://help.github.com/articles/creating-a-pull-request/).

   If your change is for fixing an issue, please write ["Fixes <issue-URL>"](https://help.github.com/articles/closing-issues-using-keywords/) in the description section of your pull request.  Github would close the issue when the owners merge your pull request.

   Please remember to specify some reviewers for your pull request.  If you don't know who are the right ones, please follow Github's recommendation.


1. Delete local and remote branches

   To keep your local workspace and your fork clean, you might want to remove merged branches:

   ```bash
   git push origin :my-cool-stuff
   git checkout develop
   git pull upstream develop
   git branch -d my-cool-stuff
   ```

### Code Review

-  Please feel free to ping your reviewers by sending them the URL of your pull request via IM or email.  Please do this after your pull request passes the CI.

- Please answer reviewers' every comment.  If you are to follow the comment, please write "Done"; please give a reason otherwise.

- If you don't want your reviewers to get overwhelmed by email notifications, you might reply their comments by [in a batch](https://help.github.com/articles/reviewing-proposed-changes-in-a-pull-request/).

- Reduce the unnecessary commits.  Some developers commit often.  It is recommended to append a sequence of small changes into one commit by running `git commit --amend` instead of `git commit`.


## Coding Standard

### Code Style

Our C/C++ code follows the [Google style guide](http://google.github.io/styleguide/cppguide.html).

Our Python code follows the [PEP8 style guide](https://www.python.org/dev/peps/pep-0008/).

Our build process helps to check the code style. 

### Unit Tests

Please remember to add related unit tests as part of your contribution.

- For C/C++ code, please follow [`google-test` Primer](https://github.com/google/googletest/blob/master/googletest/docs/primer.md) .

- For Python code, please use [Python's standard `unittest` package](http://pythontesting.net/framework/unittest/unittest-introduction/).


## Vulnerability reporting (security issues)
If you find a significant vulnerability, or evidence of one, please send an email to the security contacts that you have such information, and we'll tell you the next steps. For now, the security contacts are: DongDaxiang dongdaxiang@baidu.com, GongWeibao gongweibao@baidu.com, and WangXi wangxi16@baidu.com

Please use an email system (like Gmail) that supports hop-to-hop encryption using STARTTLS when reporting vulnerabilities. Examples of such systems include Gmail, Outlook.com, and runbox.com. See STARTTLS Everywhere if you wish to learn more about efforts to encourage the use of STARTTLS. Your email client should use encryption to communicate with your email system (i.e., if you use a web-based email client then use HTTPS, and if you use email client software then configure it to use encryption). Hop-to-hop encryption isn't as strong as end-to-end encryption, but we've decided that it's strong enough for this purpose and it's much easier to get everyone to use it.

We will gladly give credit to anyone who reports a vulnerability so that we can fix it. If you want to remain anonymous or pseudonymous instead, please let us know that; we will gladly respect your wishes.
