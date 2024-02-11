package main

func SwapKeysAndValues(m map[string]string) map[string]string {
	m_n := make(map[string]string)
	for key, val := range m {
		m_n[val] = key
	}
	return m_n
}
