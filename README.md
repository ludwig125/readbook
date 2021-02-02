```
go get github.com/sclevine/agouti
go get golang.org/x/crypto/ssh/terminal
```

```
bash-3.2$ brew install chromedriver
Error: No available formula with the name "chromedriver"
It was migrated from homebrew/core to homebrew/cask.
You can access it again by running:
  brew tap homebrew/cask
And then you can install it by running:
  brew cask install chromedriver
bash-3.2$

bash-3.2$ brew cask install chromedriver
Updating Homebrew...
==> Satisfying dependencies
==> Downloading https://chromedriver.storage.googleapis.com/2.45/chromedriver_mac64.zip
######################################################################## 100.0%
==> Verifying SHA-256 checksum for Cask 'chromedriver'.
==> Installing Cask chromedriver
==> Linking Binary 'chromedriver' to '/usr/local/bin/chromedriver'.
�  chromedriver was successfully installed!
bash-3.2$
```


WSLで構築したとき
```
$wget https://dl.google.com/linux/linux_signing_key.pub

$sudo apt-key add linux_signing_key.pub
[sudo] ludwig125 のパスワード:
OK

$echo 'deb [arch=amd64] http://dl.google.com/linux/chrome/deb/ stable main' | sudo tee /etc/apt/sources.list.d/google-chrome.list
deb [arch=amd64] http://dl.google.com/linux/chrome/deb/ stable main

$sudo apt-get update

$sudo apt-get install google-chrome-stable
略

$google-chrome --version
Google Chrome 83.0.4103.97

$wget https://chromedriver.storage.googleapis.com/83.0.4103.39/chromedriver_linux64.zip
略

$unzip chromedriver_linux64.zip
Archive:  chromedriver_linux64.zip
  inflating: chromedriver
$echo$PATH
/mnt/c/wsl/scala-2.12.5/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/usr/games:/usr/local/games:以下略
$ls
chromedriver*  chromedriver_linux64.zip*  go/  linux_signing_key.pub*

$sudo mv chromedriver /usr/local/bin

$which chromedriver
/usr/local/bin/chromedriver

$chromedriver --version
ChromeDriver 83.0.4103.39 (ccbf011cb2d2b19b506d844400483861342c20cd-refs/branch-heads/4103@{#416})
$
```