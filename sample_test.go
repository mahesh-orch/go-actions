package main

func TestSampleFunc (t *testing.T) {
  actual := sampleFunc()
  expected := "Hello Universe. You are infinite..!"
  assert.Equal(t, actual, expected)  
}
