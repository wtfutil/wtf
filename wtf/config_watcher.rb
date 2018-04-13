package wtf

import(
"time"
)

type ConfigWatcher struct {
  UpdatedAt *time.Time
}

func (watcher *ConfigWatcher) Watch() {

}
