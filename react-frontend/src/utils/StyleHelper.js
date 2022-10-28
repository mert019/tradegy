export function getOrderListStatusBadgeStyle(orderStatus) {

    let badgeStatus = '';

    switch(orderStatus) {
        case "Open":
          badgeStatus = 'status-badge-neutral';
          break;
          case "CanceledByUser":
          badgeStatus = 'status-badge-fail';
          break;
          case "CanceledBySystem":
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
