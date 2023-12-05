import {React} from 'react'
import './main.css'

const Main = () => {
   return (
      <div className="upload-container">
         <h1 className="title">Files Converter</h1>
         <div className="upload-an-ddownload">
            <label
               htmlFor="images"
               class="drop-container"
               id="dropcontainer"
            >
               <span class="drop-title">Drop files here</span>
               or
               <input
                  type="file"
                  id="images"
                  required
               />
            </label>

            <div className="convert-cont">
               <button className="convert-btn">Convert files</button>
            </div>

            <div className="download-cont">
               <button className="btn-download">
                  Download PDF
               </button>
            </div>
         </div>
      </div>
   )
}

export default Main;