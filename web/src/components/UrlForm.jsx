import  { useState } from 'react';

export default function UrlForm(){
    const [url, setUrl] = useState('');
    const [shortUrl, setShortUrl] = useState('');

    const handleSubmit = async (e) => {
        e.preventDefault();
        try{
        const response = await fetch('http://localhost:8080/shorten' , {
            method: "POST",
            headers: {
                'Content-Type': 'application/x-www-form-urlencoded',
            },
            body: `url=${encodeURIComponent(url)}`, //url to backend
        })
        if (!response.ok){
            throw new Error(`HTTP Error-> Status:  ${response.status}`)
        }
        const result = await response.text()
    //    console.log('result:', result)
        setShortUrl(result);
        alert("successfully shortened url")
    }catch(e){
        alert("error occurred while shortening the url")
        console.log("error while shortening url:", e)
    }
        }
    
    return(
        <form className="max-w-lg mt-10" onSubmit={handleSubmit}>
            <div className="flex flex-col justify-center items-center w-full ">
                <label className="text-xl text-neutral-300 mb-4">Put URL here:</label>

                    <div className="flex rounded-md shadow-sm ring-2 ring-inset ring-neutral-300 focus-within:ring-2 focus-within:ring-inset  focus-within:ring-sage-green sm:max-w-md">
                        <input className="block flex-1 py-1.5 pl-1 px-3 rounded-md border-0 bg-transparent text-neutral-600 placeholder-neutral-500 focus:border-teal focus:outline-sage-green focus:ring-0" placeholder= "https://hello.com/long-url" onChange={(e) => setUrl(e.target.value)}></input>
                    </div>
                    <button className="text-white border mt-8 bg-steel-blue rounded px-2 py-1" type="submit">Get your short URL</button>
                
            </div>
            {shortUrl &&
            <div> {shortUrl}</div>
           
            }
            {/* {
                url &&
                <div>long Url:{url}</div>
            } */}
        </form>
    )
}