export class Record {
    public Number: number;
    public Headline: string;
    public Description: string;
    public Links: string[];
    constructor(number: number, headline: string, description: string, links: string[]) {
        this.Number = number;
        this.Headline = headline;
        this.Description = description;
        this.Links = links;
    }
}
