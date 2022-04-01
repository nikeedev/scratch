// We use class syntax to define our extension object
// This isn't actually necessary, but it tends to look the best

class extension {
    /**
     * Scratch will call this method *once* when the extension loads.
     * This method's job is to tell Scratch things like the extension's ID, name, and what blocks it supports.
     */
    getInfo() {
        return {
            // `id` is the internal ID of the extension
            // It should never change!
            // If you choose to make an actual extension, please change this to something else.
            // Only the characters a-z and 0-9 can be used. No spaces or special characters.
            id: 'NikeeExtension',
  
            // `name` is what the user sees in the toolbox
            // It can be changed without breaking projects.
            name: '',
  
            blocks: [
                {
                    // `opcode` is the internal ID of the block
                    // It should never change!
                    // It corresponds to the class method with the same name.
                    opcode: 'hello1',
                    blockType: Scratch.BlockType.BOOLEAN,
                    text: 'Hello, world!  1'
                },
                {
                    // `opcode` is the internal ID of the block
                    // It should never change!
                    // It corresponds to the class method with the same name.
                    opcode: 'hello2',
                    blockType: Scratch.BlockType.REPORTER,
                    text: 'Hello, world!  2'
                },
                {
                    // `opcode` is the internal ID of the block
                    // It should never change!
                    // It corresponds to the class method with the same name.
                    opcode: 'hello3',
                    blockType: Scratch.BlockType.COMMAND,
                    text: 'Hello, world!  3'
                },
                {
                    // `opcode` is the internal ID of the block
                    // It should never change!
                    // It corresponds to the class method with the same name.
                    opcode: 'text',
                    blockType: Scratch.BlockType.REPORTER,
                    text: '\t π'
                }
            ]
        };
    }

    /**
    * Corresponds to `opcode: 'hello'` above
    */
    hello1() {
        // You can just return a value: any string, boolean, or number will work.
        // If you have to perform an asynchronous action like a request, just return a Promise.
        // The block will wait until the Promise resolves and return the resolved value.
        return true;
    }
    hello2() {
        // You can just return a value: any string, boolean, or number will work.
        // If you have to perform an asynchronous action like a request, just return a Promise.
        // The block will wait until the Promise resolves and return the resolved value.
        return "Hello world!  2";
    }
    hello3() {
        // You can just return a value: any string, boolean, or number will work.
        // If you have to perform an asynchronous action like a request, just return a Promise.
        // The block will wait until the Promise resolves and return the resolved value.
        return hello1();
    }

    text() {
        return "\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n hei";
    }

}
 
// Call Scratch.extensions.register to register your extension
// Make sure to register each extension exactly once
Scratch.extensions.register(new extension());


