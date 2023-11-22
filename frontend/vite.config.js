import fg from 'fast-glob';
import { defineConfig } from 'vite'

// https://vitejs.dev/config/
export default defineConfig({
    base: "/www/",
    // plugins: [
    //     {
    //         name: 'watch-external', // https://stackoverflow.com/questions/63373804/rollup-watch-include-directory/63548394#63548394
    //         async buildStart(){
    //             const files = await fg(["public/**/*"]);
    //             for(let file of files){
    //                 this.addWatchFile(file);
    //             }
    //         }
    //     }
    // ]
})
