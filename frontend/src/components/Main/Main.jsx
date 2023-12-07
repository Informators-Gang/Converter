import {React, useState} from 'react'
import axios from 'axios';
import Config from '../../Config';
import './main.css'

const Main = () => {

   const [name, setName] = useState("");
   const [selectedFile, setSelectedFile] = useState(null);
   const [resultFile, setResultFile] = useState(null)
   const [selectedFormat, setSelectedFormat] = useState('');
   const [filename, setFileName] = useState('default_name')

   // Function to handle the format selection
   const handleFormatChange = (event) => {
     setSelectedFormat(event.target.value);
   };

   const convertFile = () => {
      const formData = new FormData();
      formData.append("filename", name);
      formData.append("file", selectedFile);
      formData.append("convert_to", selectedFormat)

      axios
         .post(Config.backend_environment + "/convert", formData, {responseType: 'blob'})
         .then((res) => {
            setFileName(`converted_file.${selectedFormat}`);
            alert("File conversion successful");
            let urlToPreview = URL.createObjectURL(res.data); 
            setResultFile(urlToPreview)
         })
         .catch((err) => alert("File Upload Error: " + err));
   };

   const onDownloadClick = async () => {
   
      let alink = document.createElement('a')
      alink.href = resultFile;
      alink.download = filename
      alink.click()
   }

   return (
      <div className="upload-container">
         <h1 className="title">Files Converter</h1>
         <div className="upload-an-ddownload">
            <label
               htmlFor="images"
               className="drop-container"
               id="dropcontainer"
            >
               <span className="drop-title">Drop files here</span>
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


            <label htmlFor="formatDropdown" className="label">Choose format:</label>
            <select
               id="formatDropdown"
               onChange={handleFormatChange}
               value={selectedFormat}
               required
            >
               <option value="" disabled hidden>
                  Choose type
               </option>
               <option value="png">PNG</option>
               <option value="jpg">JPG</option>
               <option value="tiff">TIFF</option>
            </select>

            <div className="convert-cont">
               <button className="convert-btn" onClick={convertFile}>Convert file</button>
            </div>

            <div className="download-cont">
               <button onClick={onDownloadClick} className="btn-download">
                  Download File
               </button>
            </div>
         </div>
      </div>
   )
}

export default Main;