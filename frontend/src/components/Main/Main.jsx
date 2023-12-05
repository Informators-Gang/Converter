import {React, useState} from 'react'
import axios from 'axios';
import Config from '../../Config';
import './main.css'

const Main = () => {

   const [name, setName] = useState("");
   const [selectedFile, setSelectedFile] = useState(null);
   const [resultFile, setResultFile] = useState(null)

   const convertFile = () => {
      const formData = new FormData();
      formData.append("filename", name);
      formData.append("file", selectedFile);

      axios
         .post(Config.backend_environment + "/convert", formData)
         .then((res) => {
            alert("File Upload success");
            setResultFile(res.data)
         })
         .catch((err) => alert("File Upload Error: " + err));
   };

   const onDownloadClick = () => {
      const fileUrl = URL.createObjectURL(resultFile.file); 

      let alink = document.createElement('a')
      alink.href = fileUrl
      alink.download = resultFile.filename
      alink.click()
   }

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
                  onChange={(e) => {
                     let file = e.target.files[0];
                     setSelectedFile(file);
                     setName(file.name)
                  }}
                  required
               />
            </label>

            <div className="convert-cont">
               <button className="convert-btn" onClick={convertFile}>Convert files</button>
            </div>

            <div className="download-cont">
               <button onClick={onDownloadClick} className="btn-download">
                  Download PDF
               </button>
            </div>
         </div>
      </div>
   )
}

export default Main;