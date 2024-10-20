




let data= [10,20,30,40,50,60,70,80,90,100];

function deleteSpecificPosition(position){
    
        if(position>data.length-1)
            {
                console.log(`position is wrong ${position}`)
            }
            else{

                for(let i=position;i<data.length;i++)
                    {
                        data[i]=data[i+1]
                    }
                    data.length--
            }

            return data;
}




console.log("delete 2 position",deleteSpecificPosition(2));
console.log("delete 4 position",deleteSpecificPosition(4));
console.log("delete 7 position",deleteSpecificPosition(7));