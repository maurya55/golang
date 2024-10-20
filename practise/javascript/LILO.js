




function enqueue(num,queue){
    
    queue.length++
    queue[queue.length-1]=num;
    return queue
}

function dequeue(queue){
    
    
            queue.length--
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


