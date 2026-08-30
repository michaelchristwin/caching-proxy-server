package middleware

// r.Get("/{path:.*}", func(w http.ResponseWriter, r *http.Request) {
// 		key := r.URL.Path
// 		cache.RLock()
// 		cached, ok := cache.data[key]
// 		cache.RUnlock()
// 		if ok {
// 			// cache hit
// 			for k, vv := range cached.header {
// 				for _, v := range vv {
// 					w.Header().Add(k, v)
// 				}
// 			}
// 			w.Header().Set("X-Cache", "HIT")
// 			w.WriteHeader(cached.status)
// 			w.Write(cached.body)
// 			return
// 		}

// 		resp, err := http.Get(origin + key)
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusBadGateway)
// 			return
// 		}
// 		defer resp.Body.Close()
// 		body, _ := io.ReadAll(resp.Body)
// 		for k, vv := range resp.Header {
// 			for _, v := range vv {
// 				w.Header().Add(k, v)
// 			}
// 		}
// 		w.Header().Set("X-Cache", "MISS")
// 		w.WriteHeader(resp.StatusCode)
// 		w.Write(body)
// 		cache.Lock()
// 		cache.data[key] = cachedResponse{
// 			status: resp.StatusCode,
// 			header: resp.Header.Clone(),
// 			body:   body,
// 		}
// 		cache.Unlock()

// 	})
