# Contributing to CONTRIBUTING.md

First off, thanks for taking the time to contribute! ❤️

All types of contributions are encouraged and valued. See the [Table of Contents](#table-of-contents) for different ways to help and details about how this project handles them. Please make sure to read the relevant section before making your contribution. It will make it a lot easier for us maintainers and smooth out the experience for all involved. The community looks forward to your contributions. 🎉

## Table of Contents

- [I Have a Question](#i-have-a-question)
- [I Want To Contribute](#i-want-to-contribute)
- [Reporting Bugs](#reporting-bugs)
- [Suggesting Enhancements](#suggesting-enhancements)
- [Your First Code Contribution](#your-first-code-contribution)
- [Improving The Documentation](#improving-the-documentation)
- [Styleguides](#styleguides)
- [Commit Messages](#commit-messages)

## I Have a Question

> If you want to ask a question, we assume that you have read the available [Documentation](https://pkg.go.dev/github.com/gopinathr143/golang-linq).

Before you ask a question, it is best to search for existing [Issues](/issues) that might help you. In case you have found a suitable issue and still need clarification, you can write your question in this issue. It is also advisable to search the internet for answers first.

If you then still feel the need to ask a question and need clarification, we recommend the following:

- Open an [Issue](/issues/new).
- Provide as much context as you can about what you're running into.
- Provide project and platform versions (os, go version, etc), depending on what seems relevant.

We will then take care of the issue as soon as possible.

## I Want To Contribute

> ### Legal Notice
>
> When contributing to this project, you must agree that you have authored 100% of the content, that you have the necessary rights to the content and that the content you contribute may be provided under the project license.

### Reporting Bugs

#### Before Submitting a Bug Report

A good bug report shouldn't leave others needing to chase you up for more information. Therefore, we ask you to investigate carefully, collect information and describe the issue in detail in your report. Please complete the following steps in advance to help us fix any potential bug as fast as possible.

- Make sure that you are using the latest version.
- Determine if your bug is really a bug and not an error on your side e.g. using incompatible environment components/versions (Make sure that you have read the [documentation](). If you are looking for support, you might want to check [this section](#i-have-a-question)).
- To see if other users have experienced (and potentially already solved) the same issue you are having, check if there is not already a bug report existing for your bug or error in the [bug tracker](issues?q=label%3Abug).
- Also make sure to search the internet (including Stack Overflow) to see if users outside of the GitHub community have discussed the issue.
- Collect information about the bug:
- Stack trace (Traceback)
- OS, Platform and Version (Windows, Linux, macOS, x86, ARM)
- Version of the interpreter, compiler, SDK, runtime environment, package manager, depending on what seems relevant.
- Possibly your input and the output
- Can you reliably reproduce the issue? And can you also reproduce it with older versions?

#### How Do I Submit a Good Bug Report?

> You must never report security related issues, vulnerabilities or bugs including sensitive information to the issue tracker, or elsewhere in public.

We use GitHub issues to track bugs and errors. If you run into an issue with the project:

- Open an [Issue](/issues/new). (Since we can't be sure at this point whether it is a bug or not, we ask you not to talk about a bug yet and not to label the issue.)
- Explain the behavior you would expect and the actual behavior.
- Please provide as much context as possible and describe the _reproduction steps_ that someone else can follow to recreate the issue on their own. This usually includes your code. For good bug reports you should isolate the problem and create a reduced test case.
- Provide the information you collected in the previous section.

### Suggesting Enhancements

This section guides you through submitting an enhancement suggestion, **including completely new features and minor improvements to existing functionality**. Following these guidelines will help maintainers and the community to understand your suggestion and find related suggestions.

#### Before Submitting an Enhancement

- Make sure that you are using the latest version.
- Read the [documentation](https://pkg.go.dev/github.com/gopinathr143/golang-linq) carefully and find out if the functionality is already covered.
- Perform a [search](/issues) to see if the enhancement has already been suggested. If it has, add a comment to the existing issue instead of opening a new one.
- Find out whether your idea fits with the scope and aims of the project. It's up to you to make a strong case to convince the project's developers of the merits of this feature. Keep in mind that we want features that will be useful to the majority of our users and not just a small subset. If you're just targeting a minority of users, consider writing an add-on/plugin library.

#### How Do I Submit a Good Enhancement Suggestion?

Enhancement suggestions are tracked as [GitHub issues](/issues).

- Use a **clear and descriptive title** for the issue to identify the suggestion.
- Provide a **step-by-step description of the suggested enhancement** in as many details as possible.
- **Describe the current behavior** and **explain which behavior you expected to see instead** and why. At this point you can also tell which alternatives do not work for you.
- **Explain why this enhancement would be useful** to most users. You may also want to point out the other projects that solved it better and which could serve as inspiration.

### Your First Code Contribution

To contribute to this project, follow these steps:

- Fork the Repository: If you haven't already, fork the project's repository to your GitHub account.

- Clone Your Fork: Clone your forked repository to your local development environment.

- Create a New Branch: Create a new branch for your changes.

- Make Changes: Make the necessary changes or improvements to the documentation. Please follow the existing style and formatting guidelines.

- Write Tests: Make sure to update existing unit tests that are affected by your changes. Write unit tests for the new features that have been introduced by your changes.

- Run tests

```bash
go test ./... -v
```

- Run linter to check the formatting of your code. If golangci-lint is not installed, install it from [here](https://golangci-lint.run/usage/install/#local-installation)

```bash
golangci-lint run
```

### Improving The Documentation

We use godoc to generate and host our documentation(https://pkg.go.dev/github.com/gopinathr143/golang-linq). If you find issues with existing documentation or have suggestions for enhancements, please create an issue/PR with necessary changes.

## Styleguides

Please follow the issue and PR templates while creating issue/PR.

### Commit Messages

Clear and consistent commit messages are important for maintaining a readable and understandable version history of the project. Please follow these guidelines when crafting commit messages:

1. Use Descriptive Messages: Write a concise and meaningful message that describes the purpose of your commit. A good commit message should answer the question, "Why is this change necessary?"<br>
   ✅ Good: "Fix a bug causing the app to crash when navigating to the settings page."<br>
   ❌ Bad: "Fix issue" or "Bugfix"

2. Start with an Action Verb: Begin your commit message with a present-tense action verb, such as "Fix," "Add," "Update," or "Remove."

3. Reference Issues: If your commit is related to a specific issue or feature request, reference it by mentioning the issue number in your commit message. This helps maintain a clear link between code changes and the reasons behind them.<br>
   ✅ Good: "Fix #123: Issue with user authentication"

### PR Title

Format of the title<br>
`<type>: <subject>`<br>
The type should always be lowercase as shown below.

Allowed `<type>` values:

- `feat` for a new feature.
- `fix` for a bug fix.
- `perf` for performance improvements.
- `docs` for changes to the documentation.
- `refactor` for refactoring production code, e.g. renaming a variable.
- `test` for adding missing tests, refactoring tests; no production code change.
- `build` for updating build configuration, development tools or other changes irrelevant to the user.
- `chore` for updates that do not apply to the above, such as dependency updates.
- `ci` for changes to CI configuration files and scripts
- `revert` for revert to a previous commit
