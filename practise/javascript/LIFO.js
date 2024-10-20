




function enqueue(num,queue){
    
    // if(queue.length==0)
    //     {
    //      queue.push(num)
    //     }
    // else{
    //     // queue.unshift(num)
    //     let currnet=num
    //     for(let i=0;i<queue.length;i++)
    //         {
    //             let a=queue[i];
    //             queue[i]=currnet
    //             currnet=a
    //         }
    //         queue[queue.length]=currnet
    // }

    queue.length++
    queue[queue.length-1]=num;
    return queue
}

function dequeue(queue){
    
    if(queue.length==0)
        {
         queue==[]
        }
    else{
        // queue.unshift(num)
        for(let i=0;i<queue.length-1;i++)
            {
                queue[i]=queue[i+1]
            }
            queue.length--
    }
    return queue
}



let queue=[];

    queue=enqueue(10,queue);
console.log("Data is adedd 10 ",queue)

queue=enqueue(20,queue);
console.log("Data is adedd 20 ",queue)

queue=enqueue(30,queue);
console.log("Data is adedd 30 ",queue)

queue=enqueue(40,queue);
console.log("Data is adedd 40 ",queue)

queue=enqueue(50,queue);
console.log("Data is adedd 50 ",queue)



queue=dequeue(queue);
console.log("Remove data ",queue)

queue=dequeue(queue);
console.log("Remove data ",queue)

queue=dequeue(queue);
console.log("Remove data ",queue)


