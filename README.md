# Pixel Art to Text

Tool to convert a pixel art image to unicode characters.

Example, using [Kenney's amazing input prompts](https://kenney-assets.itch.io/input-prompts-pixel-16):

```
$ git clone git@github.com:SirGFM/pixelart-to-text.git
$ cd pixelart-to-text/
$ go build .
$ ./pixelart-to-text ./gamepad.png

      ████████████████████████                                                                              ████████████████████████
    ████████████████████████████                                                                          ████████████████████████████
  ██████▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒██████    ████████████████████████████      ████████████████████████████    ██████▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒██████
  ████▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒████  ████████████████████████████████  ████████████████████████████████  ████▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒████
  ████▒▒▒▒▒▒▒▒▓▓▒▒▒▒▓▓▓▓▓▓▒▒▒▒████  ████▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒████  ████▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒████  ████▒▒▒▒▓▓▓▓▓▓▒▒▓▓▓▓▓▓▒▒▒▒▒▒████
  ████▒▒▒▒▒▒▒▒▓▓▒▒▒▒▒▒▓▓▒▒▒▒▒▒████  ████▒▒▒▒▒▒▒▒▒▒▓▓▒▒▒▒▓▓▓▓▒▒▒▒████  ████▒▒▓▓▓▓▓▓▒▒▓▓▓▓▒▒▒▒▒▒▒▒▒▒████  ████▒▒▒▒▓▓▒▒▓▓▒▒▒▒▓▓▒▒▒▒▒▒▒▒████
  ████░░▒▒▒▒▒▒▓▓▒▒▒▒▒▒▓▓▒▒▒▒▒▒████  ████▒▒▒▒▒▒▒▒▒▒▓▓▒▒▒▒▓▓▓▓▒▒▒▒████  ████▒▒▓▓▒▒▓▓▒▒▓▓▓▓▒▒▒▒▒▒▒▒▒▒████  ████▒▒▒▒▓▓▓▓▒▒▒▒▒▒▓▓▒▒▒▒▒▒░░████
  ████▓▓▒▒▒▒▒▒▓▓▓▓▓▓▒▒▓▓▒▒▒▒▒▒████  ████▒▒▒▒▒▒▒▒▒▒▓▓▒▒▒▒▓▓▒▒▓▓▒▒████  ████▒▒▓▓▓▓▒▒▒▒▓▓▒▒▓▓▒▒▒▒▒▒▒▒████  ████▒▒▒▒▓▓▒▒▓▓▒▒▒▒▓▓▒▒▒▒▒▒▓▓████
  ████▓▓░░▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒████  ████░░▒▒▒▒▒▒▒▒▓▓▓▓▒▒▓▓▓▓▓▓▒▒████  ████▒▒▓▓▒▒▓▓▒▒▓▓▓▓▓▓▒▒▒▒▒▒░░████  ████▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒░░▓▓████
  ██████▓▓░░▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒░░████  ████▓▓░░░░▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒████  ████▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒░░░░▓▓████  ████░░▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒░░▓▓██████
    ████▓▓▓▓░░░░░░░░░░░░░░░░░░████  ██████▓▓▓▓░░░░░░░░░░░░░░░░░░████  ████░░░░░░░░░░░░░░░░░░▓▓▓▓██████  ████░░░░░░░░░░░░░░░░░░▓▓▓▓████              ████████████
    ██████▓▓▓▓▓▓░░░░░░░░░░░░▓▓████    ████████▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓████  ████▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓████████    ████▓▓░░░░░░░░░░░░▓▓▓▓▓▓██████          ████████████████████
      ██████▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓████      ████████████████████████████  ████████████████████████████      ████▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓██████          ████████▒▒▒▒▒▒▒▒████████
        ████████▓▓▓▓▓▓▓▓▓▓▓▓██████          ██████████████████████      ██████████████████████          ██████▓▓▓▓▓▓▓▓▓▓▓▓████████            ████▒▒░░░░▒▒▒▒░░░░▒▒████
          ██████████████████████                                                                          ██████████████████████            ██████▒▒▓▓▓▓░░░░▓▓▓▓▒▒██████
              ████████████████                                                                              ████████████████                ████▒▒▒▒▒▒▓▓▓▓▓▓▓▓▒▒▒▒▒▒████
                                                                                                                                            ████▒▒▒▒▒▒░░▓▓▓▓░░▒▒▒▒▒▒████
                                                                                                                                            ████▒▒▒▒░░▓▓▓▓▓▓▓▓░░▒▒▒▒████
          ████████████████                                                                                        ████████████              ████▓▓▒▒▓▓▓▓▒▒▒▒▓▓▓▓▒▒▓▓████
        ████████████████████            ████████████████████████          ████████████████████████            ████████████████████          ██████░░░░▒▒▒▒▒▒▒▒░░░░██████
      ██████▒▒▒▒▒▒▒▒▒▒▒▒██████        ████████████████████████████      ████████████████████████████        ████████▒▒▒▒▒▒▒▒████████          ████▓▓▓▓░░░░░░░░▓▓▓▓████
      ████▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒████      ██████▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒██████  ██████▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒██████      ████▒▒░░░░▒▒▒▒░░░░▒▒████          ████████▓▓▓▓▓▓▓▓████████
      ████▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒████      ████░░▒▒▒▒▒▒▒▒▒▒▓▓▒▒▒▒▒▒▒▒░░████  ████░░▒▒▒▒▒▒▒▒▓▓▒▒▒▒▒▒▒▒▒▒░░████    ██████▒▒▓▓▓▓░░░░▓▓▓▓▒▒██████          ████████████████████
      ████▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒████      ████░░▒▒▒▒▒▒▒▒▓▓▓▓▒▒▒▒▒▒▒▒░░████  ████░░▒▒▒▒▒▒▒▒▓▓▓▓▒▒▒▒▒▒▒▒░░████    ████▒▒▒▒▒▒▓▓▓▓▓▓▓▓▒▒▒▒▒▒████              ████████████
      ████▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒████      ████░░▒▒▒▒▒▒▓▓▓▓▓▓▒▒▒▒▒▒▒▒░░████  ████░░▒▒▒▒▒▒▒▒▓▓▓▓▓▓▒▒▒▒▒▒░░████    ████▒▒▒▒▒▒▒▒▓▓▓▓▒▒▒▒▒▒▒▒████
      ████░░▒▒▒▒▒▒▒▒▒▒▒▒░░████      ████░░▒▒▒▒▒▒▒▒▓▓▓▓▒▒▒▒▒▒▒▒░░████  ████░░▒▒▒▒▒▒▒▒▓▓▓▓▒▒▒▒▒▒▒▒░░████    ████▒▒▒▒▒▒▒▒▓▓▓▓▒▒▒▒▒▒▒▒████
      ████▓▓░░░░░░░░░░░░▓▓████      ████░░▒▒▒▒▒▒▒▒▒▒▓▓▒▒▒▒▒▒▒▒░░████  ████░░▒▒▒▒▒▒▒▒▓▓▒▒▒▒▒▒▒▒▒▒░░████    ████▓▓▒▒▒▒▒▒▓▓▓▓▒▒▒▒▒▒▓▓████
      ██████▓▓▓▓▓▓▓▓▓▓▓▓██████      ██████░░░░░░░░░░░░░░░░░░░░██████  ██████░░░░░░░░░░░░░░░░░░░░██████    ██████░░░░▒▒▒▒▒▒▒▒░░░░██████                              ████████████
        ████████████████████          ████████████████████████████      ████████████████████████████        ████▓▓▓▓░░░░░░░░▓▓▓▓████                            ████████████████████
          ████▓▓▓▓▓▓▓▓████              ████████████████████████          ████████████████████████          ████████▓▓▓▓▓▓▓▓████████                          ████████▒▒▒▒▒▒▒▒████████
          ████████████████                                                                                    ████████████████████                            ████▒▒▒▒▒▒░░░░▒▒▒▒▒▒████
            ████████████                                                                                          ████████████                              ██████▒▒▒▒░░▓▓▓▓░░▒▒▒▒██████
                                                                                                                                                            ████▒▒▒▒░░▓▓▓▓▓▓▓▓░░▒▒▒▒████
                                                                                                                                                            ████▒▒▒▒▓▓▓▓░░░░▓▓▓▓▒▒▒▒████
                            ████████████                                                                                                                    ████▒▒▒▒▓▓▓▓▓▓▓▓▓▓▓▓▒▒▒▒████
                          ████████████████                                    ████████████████                                    ████████████              ████▓▓▒▒▓▓▓▓▒▒▒▒▓▓▓▓▒▒▓▓████
                          ████▒▒▒▒▒▒▒▒████                                  ████████████████████                              ████████████████████          ██████░░░░▒▒▒▒▒▒▒▒░░░░██████
                          ████▒▒▒▒▒▒▒▒████                                ██████▒▒▒▒▒▒▒▒▒▒▒▒██████                          ████████▒▒▒▒▒▒▒▒████████          ████▓▓▓▓░░░░░░░░▓▓▓▓████
                    ██████████▒▒▒▒▒▒▒▒██████████                          ████▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒████                          ████▒▒░░░░░░░░░░▒▒▒▒████          ████████▓▓▓▓▓▓▓▓████████
                  ████████████▒▒▒▒▒▒▒▒████████████                        ████▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒████                        ██████▒▒▓▓▓▓▓▓▓▓▓▓░░▒▒██████          ████████████████████
                  ████▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒████                        ████▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒████                        ████▒▒▒▒▓▓▓▓░░░░▓▓▓▓▒▒▒▒████              ████████████
                  ████▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒████                        ████▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒████                        ████▒▒▒▒▓▓▓▓▓▓▓▓▓▓░░▒▒▒▒████
                  ████░░░░░░░░▒▒▒▒▒▒▒▒░░░░░░░░████                        ████░░▒▒▒▒▒▒▒▒▒▒▒▒░░████                        ████▒▒▒▒▓▓▓▓░░░░▓▓▓▓▒▒▒▒████
                  ████▓▓▓▓▓▓▓▓▒▒▒▒▒▒▒▒▓▓▓▓▓▓▓▓████                        ████▓▓░░░░░░░░░░░░▓▓████                        ████▓▓▒▒▓▓▓▓▓▓▓▓▓▓▒▒▒▒▓▓████
                  ████████████▒▒▒▒▒▒▒▒████████████                        ██████▓▓▓▓▓▓▓▓▓▓▓▓██████                        ██████░░░░▒▒▒▒▒▒▒▒░░░░██████
                    ██████████▒▒▒▒▒▒▒▒██████████                            ████████████████████                            ████▓▓▓▓░░░░░░░░▓▓▓▓████
                          ████░░░░░░░░████                                    ████▓▓▓▓▓▓▓▓████                              ████████▓▓▓▓▓▓▓▓████████
                          ████▓▓▓▓▓▓▓▓████                                    ████████████████                                ████████████████████
                          ████████████████                                      ████████████                                      ████████████
                            ████████████

```
