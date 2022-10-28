export function getOrderListStatusBadgeStyle(orderStatus) {

  let badgeStatus = '';

  switch (orderStatus) {
    case "Open":
      badgeStatus = 'status-badge-neutral';
      break;
    case "Cancelled":
      badgeStatus = 'status-badge-fail';
      break;
    case "Executed":
      badgeStatus = 'status-badge-success';
      break;
    default:
      badgeStatus = '';
      break;
  }

  return badgeStatus;
}
