import {
  Table,
  TableBody,
  TableCaption,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";

type Invoice = {
  Invoice: string;
  Status: string;
  Method: string;
  Amount: string;
};

const invoices: Invoice[] = [
  {
    Invoice: "INV001",
    Status: "Paid",
    Method: "Credit Card",
    Amount: "$250.00",
  },
  {
    Invoice: "INV002",
    Status: "Pending",
    Method: "Bank Transfer",
    Amount: "$1,200.00",
  },
  { Invoice: "INV003", Status: "Overdue", Method: "PayPal", Amount: "$75.50" },
];

export default function Notifications() {
  return (
    <div>
      <Table>
        <TableCaption>A list of your recent invoices.</TableCaption>
        <TableHeader>
          <TableRow>
            <TableHead className="w-[100px]">Invoice</TableHead>
            <TableHead>Status</TableHead>
            <TableHead>Method</TableHead>
            <TableHead className="text-right">Amount</TableHead>
          </TableRow>
        </TableHeader>
        <TableBody>
          {invoices.map((invoice) => (
            <TableRow key={invoice.Invoice}>
              <TableCell className="font-medium">{invoice.Invoice}</TableCell>
              <TableCell>{invoice.Status}</TableCell>
              <TableCell>{invoice.Method}</TableCell>
              <TableCell className="text-right">{invoice.Amount}</TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </div>
  );
}
