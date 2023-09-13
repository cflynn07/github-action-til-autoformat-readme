# GitHub Action for Autoformatting TIL README's

[![codecov](https://codecov.io/gh/cflynn07/github-action-til-autoformat-readme/branch/master/graph/badge.svg)](https://codecov.io/gh/cflynn07/github-action-til-autoformat-readme)
[![Maintainability](https://api.codeclimate.com/v1/badges/a2d85af2b4450ba36c63/maintainability)](https://codeclimate.com/github/cflynn07/github-action-til-autoformat-readme/maintainability)
[![Tests](https://github.com/cflynn07/github-action-til-autoformat-readme/workflows/Tag%20Test%20Push/badge.svg)](https://github.com/cflynn07/github-action-til-autoformat-readme/actions?query=workflow%3A%22Tag+Test+Push%22)

A GitHub action that can be used with a TIL repo to autogenerate a README.md. 

![TIL Repo Example](./Screen_Shot_2020-04-27_at_3.44.38_PM.png)

I came across [this post (Using a self-rewriting README powered by GitHub
Actions to track TILs)][1] from [Simon Willison][3] on Hacker News and thought
it was a pretty good idea. The author talks about how he uses TILs to learn in
public and how he uses GitHub actions to automatically create a formatted
README.md summary of his TILs when he pushes code. I saw this and thought,
hey if we use GitHub actions to do this, why not make a GitHub Action?

If you have a TIL repository with TILs organized into folders by category you
can add this GitHub action to generate a nice README when you push a new TIL.

### How To Use
Add this Action to your TIL repo. Here's an example:
###### .github/workflows/build.yml
```yaml
name: Build README
on:
  push:
    branches:
    - master
    paths-ignore:
    - README.md
jobs:
  build:
    runs-on: ubuntu-latest
    steps:
    - name: Check out repo
      uses: actions/checkout@v4
      with:
        # necessary for github-action-til-autoformat-readme
        fetch-depth: 0
    - name: Autoformat README
      uses: cflynn07/github-action-til-autoformat-readme@1.2.1
      with:
        description: |
          A collection of concrete writeups of small things I learn daily while working
          and researching. My goal is to work in public. I was inspired to start this
          repository after reading Simon Wilson's [hacker new post][1], and he was
          apparently inspired by Josh Branchaud's [TIL collection][2].
        footer: |
          [1]: https://simonwillison.net/2020/Apr/20/self-rewriting-readme/
          [2]: https://github.com/jbranchaud/til
        list_most_recent: 2 # optional, lists most recent TILS below description
        date_format: "2020 Jan 15:04" # optional, must align to https://golang.org/pkg/time/#Time.Format
```

### Generated README.md example

You can see an example of a sample TIL repo with the action here:  
[cflynn07/til-autoformat-action-example](https://github.com/cflynn07/til-autoformat-action-example)

```markdown
# TIL
> Today I Learned

A collection of concrete writeups of small things I learn daily while working
and researching. My goal is to work in public. I was inspired to start this
repository after reading Simon Wilson's [hacker new post][1], and he was
apparently inspired by Josh Branchaud's [TIL collection][2].

_3 TILs and counting..._

---

### 2 most recent TILs

- [How to add a CSS border](css/how-to-add-a-border.md) - Sat Apr 25 19:39:03 2020 +0800
- [How to make a div](html/how-to-make-a-div.md) - Sat Apr 25 17:53:55 2020 +0800

### Categories

- [css](#css)
- [html](#html)
- [k8s](#k8s)

### [css](#css)
- [How to add a CSS border](css/how-to-add-a-border.md)

### [html](#html)
- [How to make a div](html/how-to-make-a-div.md)

### [k8s](#k8s)
- [how to blow up a k8s cluster](k8s/how-to-blow-up-a-cluster.md)

[1]: https://simonwillison.net/2020/Apr/20/self-rewriting-readme/
[2]: https://github.com/jbranchaud/til
```

[1]: https://news.ycombinator.com/item?id=22920437
[2]: https://simonwillison.net/2020/Apr/20/self-rewriting-readme/
[3]: https://github.com/simonw
