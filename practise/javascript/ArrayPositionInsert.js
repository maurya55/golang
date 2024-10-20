



let data= [10,20,30,40,50,60,70,80,90,100];


function ArrayPostionInsert(element,num){

    let temp=num
    for(let i=element;i<data.length;i++)
        {
           let a=data[i];
           data[i]=temp;
           temp=a
        }
        data.push(temp)

        return data

}





console.log("Add 2 position 22 ",ArrayPostionInsert(2,22));
console.log("Add 4 position 44",ArrayPostionInsert(4,44));
console.log("Add 7 position 77",ArrayPostionInsert(7,77));
