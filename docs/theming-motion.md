# Motion Tokens

To keep transitions consistent, tokens include **Motion**:

```go
type Motion struct {
  Fast, Normal, Slow time.Duration
}
```

Defaults: 140ms / 180ms / 240ms. Use with `pkg/anim`:

```go
a := anim.New(anim.Config{Duration: th.Tokens.Motion.Normal, FPS: 30, Easing: anim.EaseOutCubic})
```
