

let data = [34,3,1,25,54,12,89,5,76,789];

console.log("duble loop sorting array ", data);

 // duble loop sorting array 
    for(let i=0;i<data.length;i++)
        for(let j=0;j<data.length;j++)
            {
                if(data[j]>data[j+1])
                    {
                        let a=data[j];
                            data[j]=data[j+1];
                            data[j+1]=a
                    }
            }

            console.log(data);

//  single loop sorting array

         data = [34,3,1,25,54,12,89,5,76,789];

console.log("single loop sorting array ", data);

         for(let i=0;i<data.length;i++)
            {
                if(data[i]>data[i+1])
                    {
                        let a=data[i];
                        data[i]=data[i+1]
                        data[i+1]=a

                        i=-1
                    }


            }



            console.log(data)