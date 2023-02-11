# Unit Tests

## Status

Same code but with Unit Tests

## Content

- If you are not familiar with unit tests in Go, you need to check how to write unit tests (see for example https://www.pullrequest.com/blog/unit-testing-in-go/ )
- Use the github.com/stretchr/testify/assert to test function results
- To track the code coverage, you can use

  ```
  go test -coverprofile=coverage.out ./pkg/...
  go tool cover -html=coverage.out
  ```

## Code explanation

if you checkout step3 branch, you need to understand the code:

### unit test and asserts

This is a standard way to write unit tests:

```
func TestFeature(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
    ...
		assert.Nil(t, err)
		assert.Equal(t, "foo", todo.Title)
		assert.Equal(t, "bar", todo.Body)
	})

	t.Run(not happy path", func(t *testing.T) {
    ...
		assert.NotNil(t, err)
	})
}

```

### code coverage

And as mentionned above, to get the code coverage (and analyze it) you can use go built-in tools

```
go test -coverprofile=coverage.out ./pkg/...
go tool cover -html=coverage.out
```
