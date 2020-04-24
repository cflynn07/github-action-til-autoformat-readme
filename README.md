# GitHub Action for Autoformatting TIL README's

[![codecov](https://codecov.io/gh/cflynn07/github-action-til-autoformat-readme/branch/master/graph/badge.svg)](https://codecov.io/gh/cflynn07/github-action-til-autoformat-readme)
[![Maintainability](https://api.codeclimate.com/v1/badges/a2d85af2b4450ba36c63/maintainability)](https://codeclimate.com/github/cflynn07/github-action-til-autoformat-readme/maintainability)

A GitHub action that can be used with a TIL repo to autogenerate a README.md. 

I came across [this post (Using a self-rewriting README powered by GitHub
Actions to track TILs)][1] from [Simon Willison][3] on Hacker News and thought
it was a pretty good idea. The author talks about how he uses TILs to learn in
public and also how he uses GitHub actions to automatically create a formatted
README.md summary of his TILs whenever he pushes code. I saw this and thought,
hey if we use GitHub actions to do this, why not make a GitHub Action?

### How to use
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
      uses: actions/checkout@v2
      with:
        # necessary for github-action-til-autoformat-readme
        fetch-depth: 0
    - name: Autoformat README
      uses: cflynn07/github-action-til-autoformat-readme@master
      with:
        description: |
          A collection of concrete writeups of small things I learn daily while working
          and researching. My goal is to work in public. I was inspired to start this
          repository after reading Simon Wilson's [hacker new post][1], and he was
          apparently inspired by Josh Branchaud's [TIL collection][2].
          Quick test change. Another test
        footer: |
          [1]: https://simonwillison.net/2020/Apr/20/self-rewriting-readme/
          [2]: https://github.com/jbranchaud/til
```


```markdown
# TIL
> Today I Learned

A collection of concrete writeups of small things I learn daily while working
and researching. My goal is to work in public. I was inspired to start this
repository after reading Simon Wilson's [hacker new post][1], and he was
apparently inspired by Josh Branchaud's [TIL collection].

_4 TILs and counting..._

---

### Catagories

- [bang](#bang)
- [bar](#bar)
- [biz](#biz)

### [bang](#bang)
  
- [Bang1-test here](bang/bang1-test.md)
  
### [bar](#bar)
  
- [Bar2 test here](bar/bar2-test.md)
  
### [biz](#biz)
  
- [fooballs test hahhaha](biz/fooballs-test.md)
  
- [bizbabbasdf](biz/fooballs-test2.md)
  

[1]: https://simonwillison.net/2020/Apr/20/self-rewriting-readme/
[2]: https://github.com/jbranchaud/til
```

[1]: https://news.ycombinator.com/item?id=22920437
[2]: https://simonwillison.net/2020/Apr/20/self-rewriting-readme/
[3]: https://github.com/simonw
