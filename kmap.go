package kmap

import (
  "database/sql"
  "fmt"
  "strconv"
  "strings"
  "time"
)

type Map map[string]interface{}

func Make() Map {
  return make(Map)
}

// Returns the map value as a string.
func (m Map) String(name string) string {
  if m[name] == nil {
    return ""
  }

  switch t := m[name].(type) {
  case []byte, sql.RawBytes:
    return fmt.Sprintf("%s", t)
  case string:
    return t
  }

  return fmt.Sprintf("%v", m[name])
}

// Returns the map value as an integer.
func (m Map) Int(name string) int {
  i, _ := strconv.ParseInt(fmt.Sprintf("%s", m[name]), 10, 64)
  return int(i)
}

// Returns the map value as an int64.
func (m Map) Int64(name string) int64 {
  i, _ := strconv.ParseInt(fmt.Sprintf("%s", m[name]), 10, 64)
  return i
}

// Returns the map value as a float32.
func (m Map) Float32(name string) float32 {
  f, _ := strconv.ParseFloat(fmt.Sprintf("%v", m[name]), 32)
  return float32(f)
}

// Returns the map value as a float64.
func (m Map) Float(name string) float64 {
  f, _ := strconv.ParseFloat(fmt.Sprintf("%v", m[name]), 64)
  return f
}

// Returns the map value as a bool.
func (m Map) Bool(name string) bool {
  if m[name] == nil {
    return false
  }

  switch m[name].(type) {
  default:
    b := strings.ToLower(fmt.Sprintf("%v", m[name]))
    if b == "" || b == "0" || b == "false" {
      return false
    }
  }

  return true
}

// Returns the map value as a date/time.
func (m Map) Date(name, format string) time.Time {
  date := time.Time{}

  if format == "" {
    format = "2006-01-02 15:04:05"
  }

  switch m[name].(type) {
  case time.Time:
    date = m[name].(time.Time)
  case string:
    value := m[name].(string)
    d, err := time.Parse(format, value)
    if err == nil {
      date = d
    }
  }
  return date
}

// Returns the map value as a Map.
func (m Map) Map(name string) Map {
  dict := Make()

  switch m[name].(type) {
  case map[string]interface{}:
    for k, _ := range m[name].(map[string]interface{}) {
      dict[k] = m[name].(map[string]interface{})[k]
    }
  case Map:
    dict = m[name].(Map)
  }

  return dict
}

// Returns the map value as a native Go map
func (m Map) NativeMap(name string) map[string]interface{} {
  return map[string]interface{}(m)
}
