const http = require('http');
const https = require('https');

class serviceREST {

    constructor() { }

    requestSSL = async (options)  => await handlerHTTP(https)(options);

    request = async (options)  => await handlerHTTP(http)(options);
};

handlerHTTP = (module) => {

    return ((options) => {

        return new Promise((resolve, reject) => {

            const requestHTTP = module.request(options, (response) => {
        
                /* 
                 *  the response stream's (an instance of Stream) current data. See:
                 *  
                 *  https://nodejs.org/api/stream.html#stream_event_data 
                 * 
                 */
    
                let responseBody = '';
    
                //another chunk of data has been received, so append it to `str`
                response.on('data', (chunk) => {
                    
                    responseBody += chunk.toString();
                });
                
                //the whole response has been received, so we just print it out here
                response.on('end', () => {
                    
                    resolve(JSON.parse(responseBody));
                });
            });
            
            if (options.body) 
                requestHTTP.write(options.body);
            
            requestHTTP.on("error", (err) => {
                
                console.error("Connection refused!")
                throw err;
            });
            requestHTTP.end();
        });
    })
}

// SINGLETON
// we need just an instance
module.exports = new serviceREST();