<h1 align="center">AtomicGo | schedulee</h1>

<p align="center">

<a href="https://github.com/atomicgo/schedule/releases">
<img src="https://img.shields.io/github/v/release/atomicgo/schedule?style=flat-square" alt="Latest Release">
</a>

<a href="https://codecov.io/gh/atomicgo/schedule" target="_blank">
<img src="https://img.shields.io/github/workflow/status/atomicgo/schedule/Go?label=tests&style=flat-square" alt="Tests">
</a>

<a href="https://codecov.io/gh/atomicgo/schedule" target="_blank">
<img src="https://img.shields.io/codecov/c/gh/atomicgo/schedule?color=magenta&logo=codecov&style=flat-square" alt="Coverage">
</a>

<a href="https://codecov.io/gh/atomicgo/schedule">
<!-- unittestcount:start --><img src="https://img.shields.io/badge/Unit_Tests-0-magenta?style=flat-square" alt="Unit test count"><!-- unittestcount:end -->
</a>

<a href="https://github.com/atomicgo/schedule/issues">
<img src="https://img.shields.io/github/issues/atomicgo/schedule.svg?style=flat-square" alt="Issues">
</a>

<a href="https://opensource.org/licenses/MIT" target="_blank">
<img src="https://img.shields.io/badge/License-MIT-yellow.svg?style=flat-square" alt="License: MIT">
</a>

</p>

---

<p align="center">
<strong><a href="#install">Get The Module</a></strong>
|
<strong><a href="https://pkg.go.dev/atomicgo.dev/schedule#section-documentation" target="_blank">Documentation</a></strong>
|
<strong><a href="https://github.com/atomicgo/atomicgo/blob/main/CONTRIBUTING.md" target="_blank">Contributing</a></strong>
|
<strong><a href="https://github.com/atomicgo/atomicgo/blob/main/CODE_OF_CONDUCT.md" target="_blank">Code of Conduct</a></strong>
</p>

---

<p align="center">
  <img src="https://raw.githubusercontent.com/atomicgo/atomicgo/main/assets/header.png" alt="AtomicGo">
</p>

<p align="center">
<table>
<tbody>
<td align="center">
<img width="2000" height="0"><br>
  -----------------------------------------------------------------------------------------------------
<img width="2000" height="0">
</td>
</tbody>
</table>
</p>
<h3  align="center"><pre>go get atomicgo.dev/schedule</pre></h3>
<p align="center">
<table>
<tbody>
<td align="center">
<img width="2000" height="0"><br>
   -----------------------------------------------------------------------------------------------------
<img width="2000" height="0">
</td>
</tbody>
</table>
</p>

## Description

Package schedule provides a simple scheduler for Go.

It can run a function at a given time, in a given duration, or repeatedly at a
given interval.


## Usage

#### type Task

```go
type Task struct {
}
```

Task holds information about the running task and can be used to stop running
tasks.

#### func  After

```go
func After(d time.Duration, task func()) *Task
```
After executes the task after the given duration. The function is non-blocking.
If you want to wait for the task to be executed, use the Task.Wait method.

#### func  At

```go
func At(t time.Time, task func()) *Task
```
At executes the task at the given time. The function is non-blocking. If you
want to wait for the task to be executed, use the Task.Wait method.

#### func  Every

```go
func Every(interval time.Duration, task func()) *Task
```
Every executes the task in the given interval. The function is non-blocking. If
you want to wait for the task to be executed, use the Task.Wait method.

#### func (*Task) ExecutesIn

```go
func (s *Task) ExecutesIn() time.Duration
```
ExecutesIn returns the duration until the next execution.

#### func (*Task) IsActive

```go
func (s *Task) IsActive() bool
```
IsActive returns true if the scheduler is active.

#### func (*Task) NextExecutionTime

```go
func (s *Task) NextExecutionTime() time.Time
```
NextExecutionTime returns the time when the next execution will happen.

#### func (*Task) StartedAt

```go
func (s *Task) StartedAt() time.Time
```
StartedAt returns the time when the scheduler was started.

#### func (*Task) Stop

```go
func (s *Task) Stop()
```
Stop stops the scheduler.

#### func (*Task) Wait

```go
func (s *Task) Wait()
```
Wait blocks until the scheduler is stopped. After and At will stop automatically
after the task is executed.

---

> [AtomicGo.dev](https://atomicgo.dev) &nbsp;&middot;&nbsp;
> with ❤️ by [@MarvinJWendt](https://github.com/MarvinJWendt) |
> [MarvinJWendt.com](https://marvinjwendt.com)
