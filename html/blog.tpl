
{{ define "blog"  }}
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Tailwind CSS Tutorial by Themesberg</title>
    <link rel="stylesheet" href="/src/dist/output.css">
</head>
<body>
    <header>
        <nav class="flex items-center justify-between container max-w-3xl mx-auto pb-12 pt-6">
            <!-- Logo -->
            <div class="relative"><img src="/static/images/osousa.png" class="" alt="logo"></div>

            <!-- Menu items
            <div class="text-lg text-gray-600 hidden lg:flex">
                <a href="#" class="block mt-4 lg:inline-block text-slate-600 lg:mt-0 mr-10">
                     <span
                        class="
                          bg-slate-800
                          bg-no-repeat [background-position:0_88%]
                          [background-size:100%_0em]
                          motion-safe:transition-all motion-safe:duration-200
                          hover:[background-size:100%_100%]
                          hover:text-slate-100
                          focus:[background-size:100%_100%]
                        "
                        >Home</span>
                </a>
                <a href="#" class="hover-link block mt-4 lg:inline-block hover:text-slate-600 lg:mt-0 mr-10">
                    Services
                </a>
                <a href="#" class="block mt-4 lg:inline-block hover:text-gray-700 lg:mt-0 mr-10">
                    Portfolio
                </a>
                <a href="#" class="block hover:text-gray-700 mt-4 lg:inline-block lg:mt-0 mr-10">
                    Company
                </a>
                <a href="#" class="block hover:text-gray-700 mt-4 lg:inline-block lg:mt-0">
                    Contact
                </a>
            </div>
            -->

            <div class="hidden lg:block">
      <ul class="flex items-center">
        
        <li class="group relative mr-6 mb-1">
          <div class="absolute left-0 bottom-0 z-20 h-0 w-full opacity-75 transition-all group-hover:h-6 group-hover:w-full group-hover:bg-slate-800"></div>
          <a href="/" class="relative z-30 block px-2 font-body text-lg font-medium text-slate-900 transition-colors group-hover:text-slate-200 dark:text-slate-600 dark:group-hover:text-secondary">
           Home 
          </a>
        </li>
        <li class="group relative mr-6 mb-1">
          <div class="absolute left-0 bottom-0 z-20 h-0 w-full opacity-75 transition-all group-hover:h-6 group-hover:w-full group-hover:bg-slate-800"></div>
          <a href="/" class="relative z-30 block px-2 font-body text-lg font-medium text-slate-900 transition-colors group-hover:text-slate-200 dark:text-slate-600 dark:group-hover:text-secondary">
            About
          </a>
        </li>
        <li class="group relative mr-6 mb-1">
          <div class="absolute left-0 bottom-0 z-20 h-0 w-full opacity-75 transition-all group-hover:h-6 group-hover:w-full group-hover:bg-slate-800"></div>
          <a href="/" class="relative z-30 block px-2 font-body text-lg font-medium text-slate-900 transition-colors group-hover:text-slate-200 dark:text-slate-600 dark:group-hover:text-secondary">
            Blog
          </a>
        </li>
        <li class="group relative mr-6 mb-1">
          <div class="absolute left-0 bottom-0 z-20 h-0 w-full opacity-75 transition-all group-hover:h-6 group-hover:w-full group-hover:bg-slate-800"></div>
          <a href="/" class="relative z-30 block px-2 font-body text-lg font-medium text-slate-900 transition-colors group-hover:text-slate-200 dark:text-slate-600 dark:group-hover:text-secondary">
            Contact
          </a>
        </li>

        <li>
          <i class="bx cursor-pointer text-3xl text-slate-400 dark:text-white bxs-sun" @click="themeSwitch()" :class="isDarkMode ? 'bxs-sun' : 'bxs-moon'"></i>
        </li>
      </ul>
    </div>

            
        </nav>
    </header>

    

    <div class="container max-w-3xl mx-auto">
        <div class="container flex max-w-3xl mx-auto m-5">
            <div class="relative">
                <img src="/static/images/blog.png" class="object-scale-down h-12 w-12">
            </div>
            <h1 class="ml-5 text-2xl lg:text-2xl font-normal text-slate-600 mb-2 lg:mb-6"> Latest blog entry</h1>
        </div>

        {{ range $i := .Post}}
        <article class="relative flex flex-col max-w-3xl lg:ml-auto border-b border-gray-300 pb-5 lg:max-w-3xl 2xl:max-w-3xl xl:max-w-none xl:w-[50rem]">
           <h3 class="mb-4 text-xl text-slate-900 tracking-tight font-bold dark:text-slate-400"><a href="/blog/headless-ui-v1-5">{{$i.Title}}</a></h3>
           <div class="mb-1 prose prose-slate  dark:prose-dark">
              <p class="text-md lg:text-base font-light text-gray-800 text-justify mb-8">We just released Headless UI v1.5, which includes a brand new <code>Combobox</code> component. Comboboxes are like select controls but with autocomplete/typeahead functionality, and are a great alternative to a regular select when you’re working with large datasets and want to quickly filter for the right option.</p>
           </div>
           <div class="mt-auto items-center justify-end">
              <dl>
                 <dt class="sr-only">Date</dt>
                 <dd class="text-sm leading-6 dark:text-slate-400 lg:absolute lg:top-0 lg:right-full lg:mr-8 lg:whitespace-nowrap"><time datetime="2022-02-24T12:00:00.000Z">Febuary 24, 2022</time></dd>
              </dl>
              <svg width="2" height="2" fill="currentColor" class="mx-4 text-slate-700 lg:hidden">
                 <circle cx="1" cy="1" r="1"></circle>
              </svg>
              <div class="relative float-right">
              <a class="group inline-flex items-center h-9 rounded-full text-sm font-semibold whitespace-nowrap px-3 focus:outline-none focus:ring-2 bg-slate-100 text-slate-700 hover:bg-slate-200 hover:text-slate-900 focus:ring-slate-500 dark:bg-slate-700 dark:text-slate-100 dark:hover:bg-slate-600 dark:hover:text-white dark:focus:ring-slate-500" href="/blog/headless-ui-v1-5">
                 read more
                 <span class="sr-only">
                    , <!-- -->Headless UI v1.5: The One With Comboboxes
                 </span>
                 <svg class="overflow-visible ml-3 text-slate-300 group-hover:text-slate-400 dark:text-slate-500 dark:group-hover:text-slate-400" width="3" height="6" viewBox="0 0 3 6" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                    <path d="M0 0L3 3L0 6"></path>
                 </svg>
              </a>
            </div>
           </div>
        </article> 
        {{end}}
    </div>

    <!-- Footer -->
    <footer class="container text-center fixed bottom-0 max-w-full mx-auto items-center text-white" style="background-color: #f1f1f1;">
      <div class="container text-center items-center mx-auto pt-9">
        <div class="flex justify-center mx-auto mb-9">
          <a href="#!" class="mr-9 text-gray-800">
            <svg
              aria-hidden="true"
              focusable="false"
              data-prefix="fab"
              data-icon="facebook-f"
              class="svg-inline--fa fa-facebook-f w-2.5"
              role="img"
              xmlns="http://www.w3.org/2000/svg"
              viewBox="0 0 320 512"
            >
              <path
                fill="currentColor"
                d="M279.14 288l14.22-92.66h-88.91v-60.13c0-25.35 12.42-50.06 52.24-50.06h40.42V6.26S260.43 0 225.36 0c-73.22 0-121.08 44.38-121.08 124.72v70.62H22.89V288h81.39v224h100.17V288z"
              ></path>
            </svg>
          </a>
          <a href="#!" class="mr-9 text-gray-800">
            <svg
              aria-hidden="true"
              focusable="false"
              data-prefix="fab"
              data-icon="twitter"
              class="svg-inline--fa fa-twitter w-4"
              role="img"
              xmlns="http://www.w3.org/2000/svg"
              viewBox="0 0 512 512"
            >
              <path
                fill="currentColor"
                d="M459.37 151.716c.325 4.548.325 9.097.325 13.645 0 138.72-105.583 298.558-298.558 298.558-59.452 0-114.68-17.219-161.137-47.106 8.447.974 16.568 1.299 25.34 1.299 49.055 0 94.213-16.568 130.274-44.832-46.132-.975-84.792-31.188-98.112-72.772 6.498.974 12.995 1.624 19.818 1.624 9.421 0 18.843-1.3 27.614-3.573-48.081-9.747-84.143-51.98-84.143-102.985v-1.299c13.969 7.797 30.214 12.67 47.431 13.319-28.264-18.843-46.781-51.005-46.781-87.391 0-19.492 5.197-37.36 14.294-52.954 51.655 63.675 129.3 105.258 216.365 109.807-1.624-7.797-2.599-15.918-2.599-24.04 0-57.828 46.782-104.934 104.934-104.934 30.213 0 57.502 12.67 76.67 33.137 23.715-4.548 46.456-13.32 66.599-25.34-7.798 24.366-24.366 44.833-46.132 57.827 21.117-2.273 41.584-8.122 60.426-16.243-14.292 20.791-32.161 39.308-52.628 54.253z"
              ></path>
            </svg>
          </a>
          <a href="#!" class="mr-9 text-gray-800">
            <svg
              aria-hidden="true"
              focusable="false"
              data-prefix="fab"
              data-icon="google"
              class="svg-inline--fa fa-google w-3.5"
              role="img"
              xmlns="http://www.w3.org/2000/svg"
              viewBox="0 0 488 512"
            >
              <path
                fill="currentColor"
                d="M488 261.8C488 403.3 391.1 504 248 504 110.8 504 0 393.2 0 256S110.8 8 248 8c66.8 0 123 24.5 166.3 64.9l-67.5 64.9C258.5 52.6 94.3 116.6 94.3 256c0 86.5 69.1 156.6 153.7 156.6 98.2 0 135-70.4 140.8-106.9H248v-85.3h236.1c2.3 12.7 3.9 24.9 3.9 41.4z"
              ></path>
            </svg>
          </a>
          <a href="#!" class="mr-9 text-gray-800">
            <svg
              aria-hidden="true"
              focusable="false"
              data-prefix="fab"
              data-icon="instagram"
              class="svg-inline--fa fa-instagram w-3.5"
              role="img"
              xmlns="http://www.w3.org/2000/svg"
              viewBox="0 0 448 512"
            >
              <path
                fill="currentColor"
                d="M224.1 141c-63.6 0-114.9 51.3-114.9 114.9s51.3 114.9 114.9 114.9S339 319.5 339 255.9 287.7 141 224.1 141zm0 189.6c-41.1 0-74.7-33.5-74.7-74.7s33.5-74.7 74.7-74.7 74.7 33.5 74.7 74.7-33.6 74.7-74.7 74.7zm146.4-194.3c0 14.9-12 26.8-26.8 26.8-14.9 0-26.8-12-26.8-26.8s12-26.8 26.8-26.8 26.8 12 26.8 26.8zm76.1 27.2c-1.7-35.9-9.9-67.7-36.2-93.9-26.2-26.2-58-34.4-93.9-36.2-37-2.1-147.9-2.1-184.9 0-35.8 1.7-67.6 9.9-93.9 36.1s-34.4 58-36.2 93.9c-2.1 37-2.1 147.9 0 184.9 1.7 35.9 9.9 67.7 36.2 93.9s58 34.4 93.9 36.2c37 2.1 147.9 2.1 184.9 0 35.9-1.7 67.7-9.9 93.9-36.2 26.2-26.2 34.4-58 36.2-93.9 2.1-37 2.1-147.8 0-184.8zM398.8 388c-7.8 19.6-22.9 34.7-42.6 42.6-29.5 11.7-99.5 9-132.1 9s-102.7 2.6-132.1-9c-19.6-7.8-34.7-22.9-42.6-42.6-11.7-29.5-9-99.5-9-132.1s-2.6-102.7 9-132.1c7.8-19.6 22.9-34.7 42.6-42.6 29.5-11.7 99.5-9 132.1-9s102.7-2.6 132.1 9c19.6 7.8 34.7 22.9 42.6 42.6 11.7 29.5 9 99.5 9 132.1s2.7 102.7-9 132.1z"
              ></path>
            </svg>
          </a>
          <a href="#!" class="mr-9 text-gray-800">
            <svg
              aria-hidden="true"
              focusable="false"
              data-prefix="fab"
              data-icon="linkedin-in"
              class="svg-inline--fa fa-linkedin-in w-3.5"
              role="img"
              xmlns="http://www.w3.org/2000/svg"
              viewBox="0 0 448 512"
            >
              <path
                fill="currentColor"
                d="M100.28 448H7.4V148.9h92.88zM53.79 108.1C24.09 108.1 0 83.5 0 53.8a53.79 53.79 0 0 1 107.58 0c0 29.7-24.1 54.3-53.79 54.3zM447.9 448h-92.68V302.4c0-34.7-.7-79.2-48.29-79.2-48.29 0-55.69 37.7-55.69 76.7V448h-92.78V148.9h89.08v40.8h1.3c12.4-23.5 42.69-48.3 87.88-48.3 94 0 111.28 61.9 111.28 142.3V448z"
              ></path>
            </svg>
          </a>
          <a href="#!" class="text-gray-800">
            <svg
              aria-hidden="true"
              focusable="false"
              data-prefix="fab"
              data-icon="github"
              class="svg-inline--fa fa-github w-4"
              role="img"
              xmlns="http://www.w3.org/2000/svg"
              viewBox="0 0 496 512"
            >
              <path
                fill="currentColor"
                d="M165.9 397.4c0 2-2.3 3.6-5.2 3.6-3.3.3-5.6-1.3-5.6-3.6 0-2 2.3-3.6 5.2-3.6 3-.3 5.6 1.3 5.6 3.6zm-31.1-4.5c-.7 2 1.3 4.3 4.3 4.9 2.6 1 5.6 0 6.2-2s-1.3-4.3-4.3-5.2c-2.6-.7-5.5.3-6.2 2.3zm44.2-1.7c-2.9.7-4.9 2.6-4.6 4.9.3 2 2.9 3.3 5.9 2.6 2.9-.7 4.9-2.6 4.6-4.6-.3-1.9-3-3.2-5.9-2.9zM244.8 8C106.1 8 0 113.3 0 252c0 110.9 69.8 205.8 169.5 239.2 12.8 2.3 17.3-5.6 17.3-12.1 0-6.2-.3-40.4-.3-61.4 0 0-70 15-84.7-29.8 0 0-11.4-29.1-27.8-36.6 0 0-22.9-15.7 1.6-15.4 0 0 24.9 2 38.6 25.8 21.9 38.6 58.6 27.5 72.9 20.9 2.3-16 8.8-27.1 16-33.7-55.9-6.2-112.3-14.3-112.3-110.5 0-27.5 7.6-41.3 23.6-58.9-2.6-6.5-11.1-33.3 2.6-67.9 20.9-6.5 69 27 69 27 20-5.6 41.5-8.5 62.8-8.5s42.8 2.9 62.8 8.5c0 0 48.1-33.6 69-27 13.7 34.7 5.2 61.4 2.6 67.9 16 17.7 25.8 31.5 25.8 58.9 0 96.5-58.9 104.2-114.8 110.5 9.2 7.9 17 22.9 17 46.4 0 33.7-.3 75.4-.3 83.6 0 6.5 4.6 14.4 17.3 12.1C428.2 457.8 496 362.9 496 252 496 113.3 383.5 8 244.8 8zM97.2 352.9c-1.3 1-1 3.3.7 5.2 1.6 1.6 3.9 2.3 5.2 1 1.3-1 1-3.3-.7-5.2-1.6-1.6-3.9-2.3-5.2-1zm-10.8-8.1c-.7 1.3.3 2.9 2.3 3.9 1.6 1 3.6.7 4.3-.7.7-1.3-.3-2.9-2.3-3.9-2-.6-3.6-.3-4.3.7zm32.4 35.6c-1.6 1.3-1 4.3 1.3 6.2 2.3 2.3 5.2 2.6 6.5 1 1.3-1.3.7-4.3-1.3-6.2-2.2-2.3-5.2-2.6-6.5-1zm-11.4-14.7c-1.6 1-1.6 3.6 0 5.9 1.6 2.3 4.3 3.3 5.6 2.3 1.6-1.3 1.6-3.9 0-6.2-1.4-2.3-4-3.3-5.6-2z"
              ></path>
            </svg>
          </a>
        </div>
      </div>
    
      <div class="text-center text-gray-700 p-4" style="background-color: rgba(0, 0, 0, 0.2);">
        © 2022 zero copyrights :
        <a class="text-gray-800" href="https://osousa.me/">osousa.me</a>
      </div>
    </footer>
</body>
</html>


<!--
<h3>
{{ range $i := .Post}}
{{$i.Id}}<br/>
{{$i.Title}}<br/>
{{$i.Body}}<br/>
{{$i.Date}}<br/>
{{end}}
</h3>


<ul>
    {{range index .post.lst}}
        {{if .Done}}
            <li class="done">{{.Title}}</li>
        {{else}}
            <li>{{.Title}}</li>
        {{end}}
    {{end}}
</ul>
-->

{{ end }}
