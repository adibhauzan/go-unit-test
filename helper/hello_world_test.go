package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "Adib",
			request: "Adib",
		},
		{
			name:    "Hauzan",
			request: "Hauzan",
		},
		{
			name:    "Sofyan",
			request: "Sofyan",
		},
		{
			name:    "Adib Hauzan Sofyan",
			request: "Adib Hauzan Sofyan",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}
}

func BenchmarkSub(b *testing.B) {
	b.Run("Adib", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Adib")
		}
	})
	b.Run("Hauzan", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Hauzan")
		}
	})
}

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Adib")
	}
}

func BenchmarkHelloWorldHauzan(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Hauzan")
	}
}

func TestHelloWorldTable(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{name: "HelloWorld(Adib)",
			request:  "Adib",
			expected: "Hello Adib",
		},
		{
			name:     "HelloWorld(Hauzan)",
			request:  "Hauzan",
			expected: "Hello Hauzan",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			assert.Equal(t, test.expected, result)
		})
	}
}

func TestSubTest(t *testing.T) {
	t.Run("Adib", func(t *testing.T) {
		result := HelloWorld("Adib")
		require.Equal(t, "Hello Adib", result, "Result must be Hello Adib")
		fmt.Println("TestHelloWorld With Assert Done")
	})
	t.Run("Hauzan", func(t *testing.T) {
		result := HelloWorld("Hauzan")
		require.Equal(t, "Hello Hauzan", result, "Result must be Hello Hauzan")
		fmt.Println("TestHelloWorld With Assert Done")
	})

	t.Run("Sofyan", func(t *testing.T) {
		result := HelloWorld("Sofyan")
		require.Equal(t, "Hello Sofyan", result, "Result must be Hello Sofyan")
		fmt.Println("TestHelloWorld With Assert Done")
	})
}

func TestMain(m *testing.M) {
	// Before
	fmt.Println("Before Unit Test")

	m.Run()

	// After
	fmt.Println("After Unit Test")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("Can not run on Windows OS")
	}
	result := HelloWorld("Adib")
	require.Equal(t, "Hello Adib", result, "Reusl must be Hello Adib")
}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Adib")
	if result != "Hello Adib" {

		t.Fatal("Result Programm Salah")
	}

}
func TestHelloWorldHauzan(t *testing.T) {
	result := HelloWorld("Hauzan")
	if result != "Hello Hauzan" {

		t.Fatal("Result Programm Salah")
	}
}
func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Adib")
	assert.Equal(t, "Hello Adib", result, "Reusl must be Hello Adib")
	fmt.Println("TestHelloWorld With Assert Done")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Adib")
	require.Equal(t, "Hello Adib", result, "Reusl must be Hello Adib")
	fmt.Println("TestHelloWorld With Assert Done")
}
