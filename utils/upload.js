import { create } from '@web3-storage/w3up-client';
import fs from 'fs';
import path from 'path';
import { fileURLToPath } from 'url';
import { argv } from 'process';


const spaceIdFilePath = '/home/loopassembly/Documents/hack4bengal-backend/space_id.txt';

async function uploadFile(filePath) {
    // upload done
  
    const client = await create();

   
    const account = await client.login('ashutoshanand2560@gmail.com');

    let spaceDid;

  
    if (fs.existsSync(spaceIdFilePath)) {
        
        spaceDid = fs.readFileSync(spaceIdFilePath, 'utf-8');
       
    } else {
        
        const space = await client.addSpace();
        await account.provision(space.did());
        await space.save();
        spaceDid = space.did();

     
        fs.writeFileSync(spaceIdFilePath, spaceDid);
        
    }

    
    await client.setCurrentSpace(spaceDid);

    
    const imageFile = fs.readFileSync(filePath);

    
    const imageBlob = new Blob([imageFile]);

    
    const cid = await client.uploadFile(imageBlob);
  

    return cid;
}



// done dfdf
const __filename = fileURLToPath(import.meta.url);
const __dirname = path.dirname(__filename);

if (process.argv[1] === __filename) {
    const filePath = argv[2]; 
    if (!filePath) {
        console.error('Please provide a file path as an argument.');
        process.exit(1);
    }
    uploadFile(filePath).then(cid => {
        console.log(`CID(${cid})`); 
    }).catch(err => {
        console.error(`Error uploading file: ${err}`);
        process.exit(1);
    });
}
