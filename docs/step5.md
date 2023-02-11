# Web UI

## Status

Add a VueJS UI on top of the microservice
- To list todos
- To see/edit a todo

## Tips/Howto

You need node@16
```
brew install node@16
```

cf https://www.vuemastery.com/blog/vue-router-a-tutorial-for-vue-3/

```
mkdir browser
npm install -g @vue/cli
vue create golang-skeleton-ui
> Default ([Vue 3] babel, eslint) 
> Use Yarn
cd golang-skeleton-ui
npm i vue-router@next
npm install --save axios
vue add element-plus
❯ Fully import 
npm run serve
```

and then begin to build the UI with Element-UI components (for VueJS3): https://element-plus.org/en-US/


## Code explanation

if you checkout step5 branch, you need to understand the code:

### VueJS

I am using VueJS 3 here.
The Makefile has been enhance with few commands:

```
make build_ui
make run_ui
```

### Axios

Axios allows VueJS to do AJAX call. To know the base url there are few files to check
- `browser/golang-skeleton-ui/.env` (environment variables used all the time)
- `browser/golang-skeleton-ui/.env.production` (environment variables used in production (when we do `npm run build`))
- `browser/golang-skeleton-ui/src/constants.js` specify the `API_URL` used by Axios lib
