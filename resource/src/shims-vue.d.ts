/* eslint-disable */
declare module 'axios'
declare module 'vue-json-viewer'
declare module '*.vue' {
  import type { DefineComponent } from 'vue'
  const component: DefineComponent<{}, {}, any>
  export default component
}
